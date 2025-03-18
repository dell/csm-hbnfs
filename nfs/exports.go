/*
Copyright © 2025 Dell Inc. or its subsidiaries. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package nfs

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
	//"time"
)

var exportsLock sync.Mutex

// generation is updated each time exportfs -r is called or the NFS service is restarted.
// It is used to avoid supurfulous updates.
var (
	generation       int64 // variable updated for each change
	syncedGeneration int64 // the last synced generations
	savedUpdates     int64 // the number of saved updates
	retrySleep       = 10 * time.Second
	waitTime         = 30 * time.Second
	exportsDir       = "/noderoot/etc/"
	exportsFile      = "exports"
	pathToExports    = exportsDir + exportsFile
)

const (
	chroot   = "chroot"
	noderoot = "/noderoot"
	exportfs = "/usr/sbin/exportfs"
)

// NOTE: There is an alternate implementation commented out here that
// tried to use exportfs for everything, but exportfs does not deal with netmasks for hosts.
// We do use exportfs to resynchronize the actual exports with /etc/exports.
// Almost all this code written by Codium.

func CheckExport(directory string) (bool, error) {
	exportsLock.Lock()
	defer exportsLock.Unlock()
	return checkExport(directory)
}

func checkExport(directory string) (bool, error) {
	file, err := os.Open(pathToExports)
	if err != nil {
		return false, fmt.Errorf("failed to open %s: %v", pathToExports, err)
	}
	defer file.Close()

	scanner := GetBufioScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, directory) {
			return true, nil
		}
	}

	if err := scanner.Err(); err != nil {
		return false, fmt.Errorf("error reading %s: %v", exportsDir, err)
	}

	return false, nil
}

// GetExport retrieves the export entry for the given directory from /noderoot/etc/exports.
func GetExport(directory string) (string, error) {
	exportsLock.Lock()
	defer exportsLock.Unlock()
	file, err := os.Open(pathToExports)
	if err != nil {
		return "", fmt.Errorf("failed to open %s: %v", exportsDir, err)
	}
	defer file.Close()

	scanner := GetBufioScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, directory) {
			return line, nil
		}
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error reading %s: %v", exportsDir, err)
	}

	return "", fmt.Errorf("no export entry found for %s", directory)
}

// Returns all exports matching a certain prefix.
func GetExports(prefix string) ([]string, error) {
	exportsLock.Lock()
	defer exportsLock.Unlock()

	file, err := os.Open(pathToExports)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var matches []string
	scanner := GetBufioScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, prefix) {
			matches = append(matches, line)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return matches, nil
}

var GetBufioScanner = func(file *os.File) *bufio.Scanner {
	return bufio.NewScanner(file)
}

// AddExport adds an export entry for the given directory to /noderoot/etc/exports.
func AddExport(directory, options string) (int64, error) {
	exportsLock.Lock()
	defer exportsLock.Unlock()
	exists, err := checkExport(directory)
	if err != nil {
		return generation, err
	}
	if exists {
		return generation, fmt.Errorf("export entry for %s already exists", directory)
	}

	file, err := os.OpenFile(pathToExports, os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		return generation, fmt.Errorf("failed to open %s: %v", exportsDir, err)
	}
	defer file.Close()

	entry := fmt.Sprintf("%s %s\n", directory, options)
	if _, err := file.WriteString(entry); err != nil {
		return generation, fmt.Errorf("failed to write to %s: %v", exportsDir, err)
	}
	log.Infof("AddExport %s %s completed", directory, options)
	generation = generation + 1
	return generation, nil
}

// DeleteExport deletes an export entry for the given directory from /noderoot/etc/exports.
func DeleteExport(directory string) (int64, error) {
	exportsLock.Lock()
	defer exportsLock.Unlock()
	file1, err := os.Open(pathToExports)
	if err != nil {
		return generation, fmt.Errorf("failed to open %s: %v", exportsDir, err)
	}

	var lines []string
	scanner := GetBufioScanner(file1)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if !strings.HasPrefix(line, directory) {
			lines = append(lines, line)
		}
	}

	if err := scanner.Err(); err != nil {
		file1.Close()
		return generation, fmt.Errorf("error reading %s: %v", exportsDir, err)
	}
	file1.Close()

	file2, err := os.OpenFile(pathToExports, os.O_TRUNC|os.O_WRONLY, 0o644)
	if err != nil {
		return generation, fmt.Errorf("failed to open %s: %v", exportsDir, err)
	}

	for _, line := range lines {
		if _, err := file2.WriteString(line + "\n"); err != nil {
			file2.Close()
			return generation, fmt.Errorf("failed to write to %s: %v", exportsDir, err)
		}
	}
	file2.Close()
	log.Infof("DeleteExport %s completed", directory)
	generation = generation + 1
	return generation, nil
}

// restartNFSMountd restarts the nfs-mountd service using systemctl.
// This is the last resort of getting the nfs service going again.
func restartNFSMountd(executor Executor) error {
	exportsLock.Lock()
	defer exportsLock.Unlock()
	log.Infof("restarting nfs-mountd")
	output, err := GetLocalExecutor().ExecuteCommand("chroot", "/noderoot", "container-systemctl", "restart", "nfs-mountd")
	if err != nil {
		return fmt.Errorf("failed to restart nfs-mountd: %v, output: %s", err, string(output))
	}

	// Wait for nfs-mountd to be up, with a timeout of 60 seconds
	timeout := time.After(waitTime)
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-timeout:
			return fmt.Errorf("timeout reached: nfs-mountd did not restart within %v", waitTime)
		case <-ticker.C:
			// Check if nfs-mountd is active
			if isNfsMountdActive() {
				generation++
				return nil
			}
		}
	}
}

// isNfsMountdActive checks if the nfs-mountd service is active
func isNfsMountdActive() bool {
	_, err := GetLocalExecutor().ExecuteCommand("chroot", "/noderoot", "container-systemctl", "is-active", "--quiet", "nfs-mountd")
	return err == nil
}

// AddExport adds a new NFS export.
// func AddExport(directory, clientOptions string) error {
// 	exportsLock.Lock()
// 	defer exportsLock.Unlock()

// 	cmd := exec.Command(exportfs, "-o", clientOptions, fmt.Sprintf(":%s", directory))
// 	//cmd := exec.Command(chroot, noderoot, "exportfs", "-o", clientOptions, fmt.Sprintf(":%s", directory))
// 	var stderr bytes.Buffer
// 	cmd.Stderr = &stderr

// 	if err := cmd.Run(); err != nil {
// 		return fmt.Errorf("failed to add export: %v, %s", err, stderr.String())
// 	}
// 	log.Infof("AddExport succeeded %s %s", directory, clientOptions)
// 	return nil
// }

// UpdateExport updates an existing NFS export.
// func UpdateExport(directory, clientOptions string) error {
// 	// Updating is essentially removing and then adding the export with new options
// 	if err := RemoveExport(directory); err != nil {
// 		return err
// 	}
// 	return AddExport(directory, clientOptions)
// }

// RemoveExport removes an existing NFS export.
// func RemoveExport(directory string) error {
// 	exportsLock.Lock()
// 	defer exportsLock.Unlock()

// 	cmd := exec.Command(exportfs, "-u", fmt.Sprintf(":%s", directory))
// 	///cmd := exec.Command(chroot, noderoot, "exportfs", "-u", fmt.Sprintf(":%s", directory))
// 	var stderr bytes.Buffer
// 	cmd.Stderr = &stderr

// 	if err := cmd.Run(); err != nil {
// 		return fmt.Errorf("failed to remove export: %v, %s", err, stderr.String())
// 	}
// 	log.Infof("RemoveExport succeeded %s", directory)
// 	return nil
// }

// // FindExport checks if a directory is exported.
// func FindExport(directory string) (bool, error) {
// 	exportsLock.Lock()
// 	defer exportsLock.Unlock()

// 	cmd := exec.Command("exportfs")
// 	var out bytes.Buffer
// 	cmd.Stdout = &out

// 	if err := cmd.Run(); err != nil {
// 		return false, fmt.Errorf("failed to list exports: %v", err)
// 	}

// 	exports := out.String()
// 	return strings.Contains(exports, directory), nil
// }

// GetExport retrieves the export details for a specific directory.
// func GetExport(directory string) (string, error) {
// 	exportsLock.Lock()
// 	defer exportsLock.Unlock()

// 	cmd := exec.Command(exportfs)
// 	//cmd := exec.Command(chroot, noderoot, "exportfs")
// 	var out bytes.Buffer
// 	cmd.Stdout = &out

// 	if err := cmd.Run(); err != nil {
// 		return "", fmt.Errorf("failed to list exports: %v", err)
// 	}

// 	exports := strings.Split(out.String(), "\n")
// 	for _, export := range exports {
// 		if strings.Contains(export, directory) {
// 			log.Infof("GetExport returned %s", export)
// 			return export, nil
// 		}
// 	}

// 	return "", fmt.Errorf("export for directory %s not found", directory)
// }

// GetAllExports returns a list of all current NFS exports.
// func GetAllExports() ([]string, error) {
// 	exportsLock.Lock()
// 	defer exportsLock.Unlock()

// 	cmd := exec.Command(exportfs)
// 	//cmd := exec.Command(chroot, noderoot, "exportfs")
// 	var out bytes.Buffer
// 	cmd.Stdout = &out

// 	if err := cmd.Run(); err != nil {
// 		return nil, fmt.Errorf("failed to list exports: %v", err)
// 	}

// 	exports := strings.Split(out.String(), "\n")
// 	log.Infof("GetAllExports returned %v", exports)
// 	return exports, nil
// }

// RestartNFSMountd doesn't actually restart the server.
// Instead it issues the exportfs -r command resync the kernel NFS with /noderoot/etc/exports.
func ResyncNFSMountd(generation int64) error {
	exportsLock.Lock()
	defer exportsLock.Unlock()
	if syncedGeneration >= generation {
		savedUpdates++
		log.Infof("savedUpdates %d", savedUpdates)
		return nil
	}
	var err error
	var output []byte
	for retries := 0; retries < 2; retries++ {
		output, err = GetLocalExecutor().ExecuteCommand(chroot, noderoot, exportfs, "-r", "-a")
		if err == nil {
			syncedGeneration = generation
			log.Infof("resyncing to %s successful %d", exportsDir, generation)
			return nil
		}
		log.Infof("failed resyncing nfs-mountd: %v, retries: %d, output: %s", err, retries, string(output))
		time.Sleep(retrySleep)
	}
	return err
}

var GetLocalExecutor = func() Executor {
	return &LocalExecutor{}
}
