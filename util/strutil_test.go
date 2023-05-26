package util

import (
	"log"
	"testing"
)

func TestSnakeCase(t *testing.T) {
	//EquipmentMonitoringProject
	//equipment_monitoring_project
	var words = "EquipmentMonitoringProject.CpuGhz"
	des := SnakeCase(words)
	log.Println(des)
}

func TestUuid(t *testing.T) {
	for i := 0; i < 10; i++ {
		log.Println(UUID())
	}
}
