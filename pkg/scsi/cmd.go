/*
Copyright 2015 The GoStor Authors All rights reserved.

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

package scsi

type SCSICommandType byte

var (
	TEST_UNIT_READY           SCSICommandType = 0x00
	REZERO_UNIT               SCSICommandType = 0x01
	REQUEST_SENSE             SCSICommandType = 0x03
	FORMAT_UNIT               SCSICommandType = 0x04
	READ_BLOCK_LIMITS         SCSICommandType = 0x05
	REASSIGN_BLOCKS           SCSICommandType = 0x07
	INITIALIZE_ELEMENT_STATUS SCSICommandType = 0x07
	READ_6                    SCSICommandType = 0x08
	WRITE_6                   SCSICommandType = 0x0a
	SEEK_6                    SCSICommandType = 0x0b
	READ_REVERSE              SCSICommandType = 0x0f
	WRITE_FILEMARKS           SCSICommandType = 0x10
	SPACE                     SCSICommandType = 0x11
	INQUIRY                   SCSICommandType = 0x12
	RECOVER_BUFFERED_DATA     SCSICommandType = 0x14
	MODE_SELECT               SCSICommandType = 0x15
	RESERVE                   SCSICommandType = 0x16
	RELEASE                   SCSICommandType = 0x17
	COPY                      SCSICommandType = 0x18
	ERASE                     SCSICommandType = 0x19
	MODE_SENSE                SCSICommandType = 0x1a
	START_STOP                SCSICommandType = 0x1b
	RECEIVE_DIAGNOSTIC        SCSICommandType = 0x1c
	SEND_DIAGNOSTIC           SCSICommandType = 0x1d
	ALLOW_MEDIUM_REMOVAL      SCSICommandType = 0x1e

	SET_WINDOW                           SCSICommandType = 0x24
	READ_CAPACITY                        SCSICommandType = 0x25
	READ_10                              SCSICommandType = 0x28
	WRITE_10                             SCSICommandType = 0x2a
	SEEK_10                              SCSICommandType = 0x2b
	POSITION_TO_ELEMENT                  SCSICommandType = 0x2b
	WRITE_VERIFY                         SCSICommandType = 0x2e
	VERIFY_10                            SCSICommandType = 0x2f
	SEARCH_HIGH                          SCSICommandType = 0x30
	SEARCH_EQUAL                         SCSICommandType = 0x31
	SEARCH_LOW                           SCSICommandType = 0x32
	SET_LIMITS                           SCSICommandType = 0x33
	PRE_FETCH_10                         SCSICommandType = 0x34
	READ_POSITION                        SCSICommandType = 0x34
	SYNCHRONIZE_CACHE                    SCSICommandType = 0x35
	LOCK_UNLOCK_CACHE                    SCSICommandType = 0x36
	READ_DEFECT_DATA                     SCSICommandType = 0x37
	INITIALIZE_ELEMENT_STATUS_WITH_RANGE SCSICommandType = 0x37
	MEDIUM_SCAN                          SCSICommandType = 0x38
	COMPARE                              SCSICommandType = 0x39
	COPY_VERIFY                          SCSICommandType = 0x3a
	WRITE_BUFFER                         SCSICommandType = 0x3b
	READ_BUFFER                          SCSICommandType = 0x3c
	UPDATE_BLOCK                         SCSICommandType = 0x3d
	READ_LONG                            SCSICommandType = 0x3e
	WRITE_LONG                           SCSICommandType = 0x3f
	CHANGE_DEFINITION                    SCSICommandType = 0x40
	WRITE_SAME                           SCSICommandType = 0x41
	UNMAP                                SCSICommandType = 0x42
	READ_TOC                             SCSICommandType = 0x43
	GET_CONFIGURATION                    SCSICommandType = 0x46
	LOG_SELECT                           SCSICommandType = 0x4c
	LOG_SENSE                            SCSICommandType = 0x4d
	READ_DISK_INFO                       SCSICommandType = 0x51
	READ_TRACK_INFO                      SCSICommandType = 0x52
	MODE_SELECT_10                       SCSICommandType = 0x55
	RESERVE_10                           SCSICommandType = 0x56
	RELEASE_10                           SCSICommandType = 0x57
	MODE_SENSE_10                        SCSICommandType = 0x5a
	CLOSE_TRACK                          SCSICommandType = 0x5b
	READ_BUFFER_CAP                      SCSICommandType = 0x5c
	PERSISTENT_RESERVE_IN                SCSICommandType = 0x5e
	PERSISTENT_RESERVE_OUT               SCSICommandType = 0x5f
	VARLEN_CDB                           SCSICommandType = 0x7f
	READ_16                              SCSICommandType = 0x88
	COMPARE_AND_WRITE                    SCSICommandType = 0x89
	WRITE_16                             SCSICommandType = 0x8a
	ORWRITE_16                           SCSICommandType = 0x8b
	WRITE_VERIFY_16                      SCSICommandType = 0x8e
	VERIFY_16                            SCSICommandType = 0x8f
	PRE_FETCH_16                         SCSICommandType = 0x90
	SYNCHRONIZE_CACHE_16                 SCSICommandType = 0x91
	WRITE_SAME_16                        SCSICommandType = 0x93
	SERVICE_ACTION_IN                    SCSICommandType = 0x9e
	SAI_READ_CAPACITY_16                 SCSICommandType = 0x10
	SAI_GET_LBA_STATUS                   SCSICommandType = 0x12
	REPORT_LUNS                          SCSICommandType = 0xa0
	MAINT_PROTOCOL_IN                    SCSICommandType = 0xa3
	MOVE_MEDIUM                          SCSICommandType = 0xa5
	EXCHANGE_MEDIUM                      SCSICommandType = 0xa6
	READ_12                              SCSICommandType = 0xa8
	WRITE_12                             SCSICommandType = 0xaa
	GET_PERFORMACE                       SCSICommandType = 0xac
	READ_DVD_STRUCTURE                   SCSICommandType = 0xad
	WRITE_VERIFY_12                      SCSICommandType = 0xae
	VERIFY_12                            SCSICommandType = 0xaf
	SEARCH_HIGH_12                       SCSICommandType = 0xb0
	SEARCH_EQUAL_12                      SCSICommandType = 0xb1
	SEARCH_LOW_12                        SCSICommandType = 0xb2
	READ_ELEMENT_STATUS                  SCSICommandType = 0xb8
	SEND_VOLUME_TAG                      SCSICommandType = 0xb6
	SET_STREAMING                        SCSICommandType = 0xb6
	SET_CD_SPEED                         SCSICommandType = 0xbb
	WRITE_LONG_2                         SCSICommandType = 0xea
)

type SCSIPRServiceAction byte
type SCSIPRType byte

var (
	/* PERSISTENT_RESERVE_IN service action codes */
	PR_IN_READ_KEYS           SCSIPRServiceAction = 0x00
	PR_IN_READ_RESERVATION    SCSIPRServiceAction = 0x01
	PR_IN_REPORT_CAPABILITIES SCSIPRServiceAction = 0x02
	PR_IN_READ_FULL_STATUS    SCSIPRServiceAction = 0x03

	/* PERSISTENT_RESERVE_OUT service action codes */
	PR_OUT_REGISTER                         SCSIPRServiceAction = 0x00
	PR_OUT_RESERVE                          SCSIPRServiceAction = 0x01
	PR_OUT_RELEASE                          SCSIPRServiceAction = 0x02
	PR_OUT_CLEAR                            SCSIPRServiceAction = 0x03
	PR_OUT_PREEMPT                          SCSIPRServiceAction = 0x04
	PR_OUT_PREEMPT_AND_ABORT                SCSIPRServiceAction = 0x05
	PR_OUT_REGISTER_AND_IGNORE_EXISTING_KEY SCSIPRServiceAction = 0x06
	PR_OUT_REGISTER_AND_MOVE                SCSIPRServiceAction = 0x07

	/* Persistent Reservation scope */
	PR_LU_SCOPE byte = 0x00

	/* Persistent Reservation Type Mask format */
	PR_TYPE_WRITE_EXCLUSIVE          SCSIPRType = 0x01
	PR_TYPE_EXCLUSIVE_ACCESS         SCSIPRType = 0x03
	PR_TYPE_WRITE_EXCLUSIVE_REGONLY  SCSIPRType = 0x05
	PR_TYPE_EXCLUSIVE_ACCESS_REGONLY SCSIPRType = 0x06
	PR_TYPE_WRITE_EXCLUSIVE_ALLREG   SCSIPRType = 0x07
	PR_TYPE_EXCLUSIVE_ACCESS_ALLREG  SCSIPRType = 0x08
)

type SCSIDataDirection int

const (
	SCSIDataNone = iota
	SCSIDataWrite
	SCSIDataRead
	SCSIDataBidirection
)

const (
	CBD_GROUPID_0 = iota
	CBD_GROUPID_1
	CBD_GROUPID_2
	CBD_GROUPID_3
	CBD_GROUPID_4
	CBD_GROUPID_5
	CBD_GROUPID_6
	CBD_GROUPID_7
)

const (
	CDB_GROUP0 = 6  /*  6-byte commands */
	CDB_GROUP1 = 10 /* 10-byte commands */
	CDB_GROUP2 = 10 /* 10-byte commands */
	CDB_GROUP3 = 0  /* reserved */
	CDB_GROUP4 = 16 /* 16-byte commands */
	CDB_GROUP5 = 12 /* 12-byte commands */
	CDB_GROUP6 = 0  /* vendor specific  */
	CDB_GROUP7 = 0  /* vendor specific  */
)

type SCSIDataBuffer struct {
	Buffer         uint64
	Length         uint64
	TransferLength uint32
	Resid          int32
}

type SCSICommand struct {
	Target       *SCSITarget
	DeviceID     uint64
	Device       *SCSILu
	State        uint64
	Direction    SCSIDataDirection
	InSDBBuffer  *SCSIDataBuffer
	OutSDBBuffer *SCSIDataBuffer
	// Command ITN ID
	CommandITNID uint64
	Offset       uint64
	TL           uint32
	SCB          *[]byte
	SCBLength    int
	Lun          []uint8
	Attribute    int
	Tag          uint64
	Result       int
	SenseBuffer  []byte
	SenseLength  int
}

func SCSICDBGroupID(opcode byte) byte {
	return ((opcode >> 5) & 0x7)
}