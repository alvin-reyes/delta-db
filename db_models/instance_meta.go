package db_models

import (
	"gorm.io/gorm"
	"time"
)

type InstanceMeta struct {
	// gorm id
	ID                               uint64    `gorm:"primary_key" json:"id"`
	MemoryLimit                      uint64    `json:"memory_limit"`
	CpuLimit                         uint64    `json:"cpu_limit"`
	StorageLimit                     uint64    `json:"storage_limit"`
	DisableRequest                   bool      `json:"disable_requests"`
	DisableCommitmentPieceGeneration bool      `json:"disable_commitment_piece_generation"`
	DisableStorageDeal               bool      `json:"disable_storage_deal"`
	DisableOnlineDeals               bool      `json:"disable_online_deals"`
	DisableOfflineDeals              bool      `json:"disable_offline_deals"`
	NumberOfCpus                     uint64    `json:"number_of_cpus"`
	SystemMemory                     uint64    `json:"system_memory"`
	HeapMemory                       uint64    `json:"heap_memory"`
	HeapInUse                        uint64    `json:"heap_in_use"`
	StackInUse                       uint64    `json:"stack_in_use"`
	InstanceStart                    time.Time `json:"instance_start"`
	BytesPerCpu                      uint64    `json:"bytes_per_cpu"`
	CreatedAt                        time.Time `json:"created_at"`
	UpdatedAt                        time.Time `json:"updated_at"`
}

func (u *InstanceMeta) BeforeSave(tx *gorm.DB) (err error) {
	return
}

func (u *InstanceMeta) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (u *InstanceMeta) AfterSave(tx *gorm.DB) (err error) {
	return
}
