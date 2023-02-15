// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.2
// source: google/maps/routing/v2/toll_passes.proto

package routing

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// List of toll passes around the world that we support.
type TollPass int32

const (
	// Not used. If this value is used, then the request fails.
	TollPass_TOLL_PASS_UNSPECIFIED TollPass = 0
	// Sydney toll pass. See additional details at https://www.myetoll.com.au.
	TollPass_AU_ETOLL_TAG TollPass = 82
	// Sydney toll pass. See additional details at https://www.tollpay.com.au.
	TollPass_AU_EWAY_TAG TollPass = 83
	// Australia-wide toll pass.
	// See additional details at https://www.linkt.com.au/.
	TollPass_AU_LINKT TollPass = 2
	// Argentina toll pass. See additional details at https://telepase.com.ar
	TollPass_AR_TELEPASE TollPass = 3
	// Brazil toll pass. See additional details at https://www.autoexpreso.com
	TollPass_BR_AUTO_EXPRESO TollPass = 81
	// Brazil toll pass. See additional details at https://conectcar.com.
	TollPass_BR_CONECTCAR TollPass = 7
	// Brazil toll pass. See additional details at https://movemais.com.
	TollPass_BR_MOVE_MAIS TollPass = 8
	// Brazil toll pass. See additional details at https://pasorapido.gob.do/
	TollPass_BR_PASSA_RAPIDO TollPass = 88
	// Brazil toll pass. See additional details at https://www.semparar.com.br.
	TollPass_BR_SEM_PARAR TollPass = 9
	// Brazil toll pass. See additional details at https://taggy.com.br.
	TollPass_BR_TAGGY TollPass = 10
	// Brazil toll pass. See additional details at
	// https://veloe.com.br/site/onde-usar.
	TollPass_BR_VELOE TollPass = 11
	// Canada to United States border crossing.
	TollPass_CA_US_AKWASASNE_SEAWAY_CORPORATE_CARD TollPass = 84
	// Canada to United States border crossing.
	TollPass_CA_US_AKWASASNE_SEAWAY_TRANSIT_CARD TollPass = 85
	// Ontario, Canada to Michigan, United States border crossing.
	TollPass_CA_US_BLUE_WATER_EDGE_PASS TollPass = 18
	// Ontario, Canada to Michigan, United States border crossing.
	TollPass_CA_US_CONNEXION TollPass = 19
	// Canada to United States border crossing.
	TollPass_CA_US_NEXUS_CARD TollPass = 20
	// Indonesia.
	// E-card provided by multiple banks used to pay for tolls. All e-cards
	// via banks are charged the same so only one enum value is needed. E.g.
	// - Bank Mandiri https://www.bankmandiri.co.id/e-money
	// - BCA https://www.bca.co.id/flazz
	// - BNI https://www.bni.co.id/id-id/ebanking/tapcash
	TollPass_ID_E_TOLL TollPass = 16
	// India.
	TollPass_IN_FASTAG TollPass = 78
	// India, HP state plate exemption.
	TollPass_IN_LOCAL_HP_PLATE_EXEMPT TollPass = 79
	// Mexico toll pass.
	// https://iave.capufe.gob.mx/#/
	TollPass_MX_IAVE TollPass = 90
	// Mexico
	// https://www.pase.com.mx
	TollPass_MX_PASE TollPass = 91
	// Mexico
	//
	//	https://operadoravial.com/quick-pass/
	TollPass_MX_QUICKPASS TollPass = 93
	// http://appsh.chihuahua.gob.mx/transparencia/?doc=/ingresos/TelepeajeFormato4.pdf
	TollPass_MX_SISTEMA_TELEPEAJE_CHIHUAHUA TollPass = 89
	// Mexico
	TollPass_MX_TAG_IAVE TollPass = 12
	// Mexico toll pass company. One of many operating in Mexico City. See
	// additional details at https://www.televia.com.mx.
	TollPass_MX_TAG_TELEVIA TollPass = 13
	// Mexico toll pass company. One of many operating in Mexico City.
	// https://www.televia.com.mx
	TollPass_MX_TELEVIA TollPass = 92
	// Mexico toll pass. See additional details at
	// https://www.viapass.com.mx/viapass/web_home.aspx.
	TollPass_MX_VIAPASS TollPass = 14
	// AL, USA.
	TollPass_US_AL_FREEDOM_PASS TollPass = 21
	// AK, USA.
	TollPass_US_AK_ANTON_ANDERSON_TUNNEL_BOOK_OF_10_TICKETS TollPass = 22
	// CA, USA.
	TollPass_US_CA_FASTRAK TollPass = 4
	// Indicates driver has any FasTrak pass in addition to the DMV issued Clean
	// Air Vehicle (CAV) sticker.
	// https://www.bayareafastrak.org/en/guide/doINeedFlex.shtml
	TollPass_US_CA_FASTRAK_CAV_STICKER TollPass = 86
	// CO, USA.
	TollPass_US_CO_EXPRESSTOLL TollPass = 23
	// CO, USA.
	TollPass_US_CO_GO_PASS TollPass = 24
	// DE, USA.
	TollPass_US_DE_EZPASSDE TollPass = 25
	// FL, USA.
	TollPass_US_FL_BOB_SIKES_TOLL_BRIDGE_PASS TollPass = 65
	// FL, USA.
	TollPass_US_FL_DUNES_COMMUNITY_DEVELOPMENT_DISTRICT_EXPRESSCARD TollPass = 66
	// FL, USA.
	TollPass_US_FL_EPASS TollPass = 67
	// FL, USA.
	TollPass_US_FL_GIBA_TOLL_PASS TollPass = 68
	// FL, USA.
	TollPass_US_FL_LEEWAY TollPass = 69
	// FL, USA.
	TollPass_US_FL_SUNPASS TollPass = 70
	// FL, USA.
	TollPass_US_FL_SUNPASS_PRO TollPass = 71
	// IL, USA.
	TollPass_US_IL_EZPASSIL TollPass = 73
	// IL, USA.
	TollPass_US_IL_IPASS TollPass = 72
	// IN, USA.
	TollPass_US_IN_EZPASSIN TollPass = 26
	// KS, USA.
	TollPass_US_KS_BESTPASS_HORIZON TollPass = 27
	// KS, USA.
	TollPass_US_KS_KTAG TollPass = 28
	// KS, USA.
	TollPass_US_KS_NATIONALPASS TollPass = 29
	// KS, USA.
	TollPass_US_KS_PREPASS_ELITEPASS TollPass = 30
	// KY, USA.
	TollPass_US_KY_RIVERLINK TollPass = 31
	// LA, USA.
	TollPass_US_LA_GEAUXPASS TollPass = 32
	// LA, USA.
	TollPass_US_LA_TOLL_TAG TollPass = 33
	// MA, USA.
	TollPass_US_MA_EZPASSMA TollPass = 6
	// MD, USA.
	TollPass_US_MD_EZPASSMD TollPass = 34
	// ME, USA.
	TollPass_US_ME_EZPASSME TollPass = 35
	// MI, USA.
	TollPass_US_MI_AMBASSADOR_BRIDGE_PREMIER_COMMUTER_CARD TollPass = 36
	// MI, USA.
	TollPass_US_MI_GROSSE_ILE_TOLL_BRIDGE_PASS_TAG TollPass = 37
	// MI, USA.
	TollPass_US_MI_IQ_PROX_CARD TollPass = 38
	// MI, USA.
	TollPass_US_MI_MACKINAC_BRIDGE_MAC_PASS TollPass = 39
	// MI, USA.
	TollPass_US_MI_NEXPRESS_TOLL TollPass = 40
	// MN, USA.
	TollPass_US_MN_EZPASSMN TollPass = 41
	// NC, USA.
	TollPass_US_NC_EZPASSNC TollPass = 42
	// NC, USA.
	TollPass_US_NC_PEACH_PASS TollPass = 87
	// NC, USA.
	TollPass_US_NC_QUICK_PASS TollPass = 43
	// NH, USA.
	TollPass_US_NH_EZPASSNH TollPass = 80
	// NJ, USA.
	TollPass_US_NJ_DOWNBEACH_EXPRESS_PASS TollPass = 75
	// NJ, USA.
	TollPass_US_NJ_EZPASSNJ TollPass = 74
	// NY, USA.
	TollPass_US_NY_EXPRESSPASS TollPass = 76
	// NY, USA.
	TollPass_US_NY_EZPASSNY TollPass = 77
	// OH, USA.
	TollPass_US_OH_EZPASSOH TollPass = 44
	// PA, USA.
	TollPass_US_PA_EZPASSPA TollPass = 45
	// RI, USA.
	TollPass_US_RI_EZPASSRI TollPass = 46
	// SC, USA.
	TollPass_US_SC_PALPASS TollPass = 47
	// TX, USA.
	TollPass_US_TX_BANCPASS TollPass = 48
	// TX, USA.
	TollPass_US_TX_DEL_RIO_PASS TollPass = 49
	// TX, USA.
	TollPass_US_TX_EFAST_PASS TollPass = 50
	// TX, USA.
	TollPass_US_TX_EAGLE_PASS_EXPRESS_CARD TollPass = 51
	// TX, USA.
	TollPass_US_TX_EPTOLL TollPass = 52
	// TX, USA.
	TollPass_US_TX_EZ_CROSS TollPass = 53
	// TX, USA.
	TollPass_US_TX_EZTAG TollPass = 54
	// TX, USA.
	TollPass_US_TX_LAREDO_TRADE_TAG TollPass = 55
	// TX, USA.
	TollPass_US_TX_PLUSPASS TollPass = 56
	// TX, USA.
	TollPass_US_TX_TOLLTAG TollPass = 57
	// TX, USA.
	TollPass_US_TX_TXTAG TollPass = 58
	// TX, USA.
	TollPass_US_TX_XPRESS_CARD TollPass = 59
	// UT, USA.
	TollPass_US_UT_ADAMS_AVE_PARKWAY_EXPRESSCARD TollPass = 60
	// VA, USA.
	TollPass_US_VA_EZPASSVA TollPass = 61
	// WA, USA.
	TollPass_US_WA_BREEZEBY TollPass = 17
	// WA, USA.
	TollPass_US_WA_GOOD_TO_GO TollPass = 1
	// WV, USA.
	TollPass_US_WV_EZPASSWV TollPass = 62
	// WV, USA.
	TollPass_US_WV_MEMORIAL_BRIDGE_TICKETS TollPass = 63
	// WV, USA.
	TollPass_US_WV_NEWELL_TOLL_BRIDGE_TICKET TollPass = 64
)

// Enum value maps for TollPass.
var (
	TollPass_name = map[int32]string{
		0:  "TOLL_PASS_UNSPECIFIED",
		82: "AU_ETOLL_TAG",
		83: "AU_EWAY_TAG",
		2:  "AU_LINKT",
		3:  "AR_TELEPASE",
		81: "BR_AUTO_EXPRESO",
		7:  "BR_CONECTCAR",
		8:  "BR_MOVE_MAIS",
		88: "BR_PASSA_RAPIDO",
		9:  "BR_SEM_PARAR",
		10: "BR_TAGGY",
		11: "BR_VELOE",
		84: "CA_US_AKWASASNE_SEAWAY_CORPORATE_CARD",
		85: "CA_US_AKWASASNE_SEAWAY_TRANSIT_CARD",
		18: "CA_US_BLUE_WATER_EDGE_PASS",
		19: "CA_US_CONNEXION",
		20: "CA_US_NEXUS_CARD",
		16: "ID_E_TOLL",
		78: "IN_FASTAG",
		79: "IN_LOCAL_HP_PLATE_EXEMPT",
		90: "MX_IAVE",
		91: "MX_PASE",
		93: "MX_QUICKPASS",
		89: "MX_SISTEMA_TELEPEAJE_CHIHUAHUA",
		12: "MX_TAG_IAVE",
		13: "MX_TAG_TELEVIA",
		92: "MX_TELEVIA",
		14: "MX_VIAPASS",
		21: "US_AL_FREEDOM_PASS",
		22: "US_AK_ANTON_ANDERSON_TUNNEL_BOOK_OF_10_TICKETS",
		4:  "US_CA_FASTRAK",
		86: "US_CA_FASTRAK_CAV_STICKER",
		23: "US_CO_EXPRESSTOLL",
		24: "US_CO_GO_PASS",
		25: "US_DE_EZPASSDE",
		65: "US_FL_BOB_SIKES_TOLL_BRIDGE_PASS",
		66: "US_FL_DUNES_COMMUNITY_DEVELOPMENT_DISTRICT_EXPRESSCARD",
		67: "US_FL_EPASS",
		68: "US_FL_GIBA_TOLL_PASS",
		69: "US_FL_LEEWAY",
		70: "US_FL_SUNPASS",
		71: "US_FL_SUNPASS_PRO",
		73: "US_IL_EZPASSIL",
		72: "US_IL_IPASS",
		26: "US_IN_EZPASSIN",
		27: "US_KS_BESTPASS_HORIZON",
		28: "US_KS_KTAG",
		29: "US_KS_NATIONALPASS",
		30: "US_KS_PREPASS_ELITEPASS",
		31: "US_KY_RIVERLINK",
		32: "US_LA_GEAUXPASS",
		33: "US_LA_TOLL_TAG",
		6:  "US_MA_EZPASSMA",
		34: "US_MD_EZPASSMD",
		35: "US_ME_EZPASSME",
		36: "US_MI_AMBASSADOR_BRIDGE_PREMIER_COMMUTER_CARD",
		37: "US_MI_GROSSE_ILE_TOLL_BRIDGE_PASS_TAG",
		38: "US_MI_IQ_PROX_CARD",
		39: "US_MI_MACKINAC_BRIDGE_MAC_PASS",
		40: "US_MI_NEXPRESS_TOLL",
		41: "US_MN_EZPASSMN",
		42: "US_NC_EZPASSNC",
		87: "US_NC_PEACH_PASS",
		43: "US_NC_QUICK_PASS",
		80: "US_NH_EZPASSNH",
		75: "US_NJ_DOWNBEACH_EXPRESS_PASS",
		74: "US_NJ_EZPASSNJ",
		76: "US_NY_EXPRESSPASS",
		77: "US_NY_EZPASSNY",
		44: "US_OH_EZPASSOH",
		45: "US_PA_EZPASSPA",
		46: "US_RI_EZPASSRI",
		47: "US_SC_PALPASS",
		48: "US_TX_BANCPASS",
		49: "US_TX_DEL_RIO_PASS",
		50: "US_TX_EFAST_PASS",
		51: "US_TX_EAGLE_PASS_EXPRESS_CARD",
		52: "US_TX_EPTOLL",
		53: "US_TX_EZ_CROSS",
		54: "US_TX_EZTAG",
		55: "US_TX_LAREDO_TRADE_TAG",
		56: "US_TX_PLUSPASS",
		57: "US_TX_TOLLTAG",
		58: "US_TX_TXTAG",
		59: "US_TX_XPRESS_CARD",
		60: "US_UT_ADAMS_AVE_PARKWAY_EXPRESSCARD",
		61: "US_VA_EZPASSVA",
		17: "US_WA_BREEZEBY",
		1:  "US_WA_GOOD_TO_GO",
		62: "US_WV_EZPASSWV",
		63: "US_WV_MEMORIAL_BRIDGE_TICKETS",
		64: "US_WV_NEWELL_TOLL_BRIDGE_TICKET",
	}
	TollPass_value = map[string]int32{
		"TOLL_PASS_UNSPECIFIED":                 0,
		"AU_ETOLL_TAG":                          82,
		"AU_EWAY_TAG":                           83,
		"AU_LINKT":                              2,
		"AR_TELEPASE":                           3,
		"BR_AUTO_EXPRESO":                       81,
		"BR_CONECTCAR":                          7,
		"BR_MOVE_MAIS":                          8,
		"BR_PASSA_RAPIDO":                       88,
		"BR_SEM_PARAR":                          9,
		"BR_TAGGY":                              10,
		"BR_VELOE":                              11,
		"CA_US_AKWASASNE_SEAWAY_CORPORATE_CARD": 84,
		"CA_US_AKWASASNE_SEAWAY_TRANSIT_CARD":   85,
		"CA_US_BLUE_WATER_EDGE_PASS":            18,
		"CA_US_CONNEXION":                       19,
		"CA_US_NEXUS_CARD":                      20,
		"ID_E_TOLL":                             16,
		"IN_FASTAG":                             78,
		"IN_LOCAL_HP_PLATE_EXEMPT":              79,
		"MX_IAVE":                               90,
		"MX_PASE":                               91,
		"MX_QUICKPASS":                          93,
		"MX_SISTEMA_TELEPEAJE_CHIHUAHUA":        89,
		"MX_TAG_IAVE":                           12,
		"MX_TAG_TELEVIA":                        13,
		"MX_TELEVIA":                            92,
		"MX_VIAPASS":                            14,
		"US_AL_FREEDOM_PASS":                    21,
		"US_AK_ANTON_ANDERSON_TUNNEL_BOOK_OF_10_TICKETS": 22,
		"US_CA_FASTRAK":                    4,
		"US_CA_FASTRAK_CAV_STICKER":        86,
		"US_CO_EXPRESSTOLL":                23,
		"US_CO_GO_PASS":                    24,
		"US_DE_EZPASSDE":                   25,
		"US_FL_BOB_SIKES_TOLL_BRIDGE_PASS": 65,
		"US_FL_DUNES_COMMUNITY_DEVELOPMENT_DISTRICT_EXPRESSCARD": 66,
		"US_FL_EPASS":             67,
		"US_FL_GIBA_TOLL_PASS":    68,
		"US_FL_LEEWAY":            69,
		"US_FL_SUNPASS":           70,
		"US_FL_SUNPASS_PRO":       71,
		"US_IL_EZPASSIL":          73,
		"US_IL_IPASS":             72,
		"US_IN_EZPASSIN":          26,
		"US_KS_BESTPASS_HORIZON":  27,
		"US_KS_KTAG":              28,
		"US_KS_NATIONALPASS":      29,
		"US_KS_PREPASS_ELITEPASS": 30,
		"US_KY_RIVERLINK":         31,
		"US_LA_GEAUXPASS":         32,
		"US_LA_TOLL_TAG":          33,
		"US_MA_EZPASSMA":          6,
		"US_MD_EZPASSMD":          34,
		"US_ME_EZPASSME":          35,
		"US_MI_AMBASSADOR_BRIDGE_PREMIER_COMMUTER_CARD": 36,
		"US_MI_GROSSE_ILE_TOLL_BRIDGE_PASS_TAG":         37,
		"US_MI_IQ_PROX_CARD":                            38,
		"US_MI_MACKINAC_BRIDGE_MAC_PASS":                39,
		"US_MI_NEXPRESS_TOLL":                           40,
		"US_MN_EZPASSMN":                                41,
		"US_NC_EZPASSNC":                                42,
		"US_NC_PEACH_PASS":                              87,
		"US_NC_QUICK_PASS":                              43,
		"US_NH_EZPASSNH":                                80,
		"US_NJ_DOWNBEACH_EXPRESS_PASS":                  75,
		"US_NJ_EZPASSNJ":                                74,
		"US_NY_EXPRESSPASS":                             76,
		"US_NY_EZPASSNY":                                77,
		"US_OH_EZPASSOH":                                44,
		"US_PA_EZPASSPA":                                45,
		"US_RI_EZPASSRI":                                46,
		"US_SC_PALPASS":                                 47,
		"US_TX_BANCPASS":                                48,
		"US_TX_DEL_RIO_PASS":                            49,
		"US_TX_EFAST_PASS":                              50,
		"US_TX_EAGLE_PASS_EXPRESS_CARD":                 51,
		"US_TX_EPTOLL":                                  52,
		"US_TX_EZ_CROSS":                                53,
		"US_TX_EZTAG":                                   54,
		"US_TX_LAREDO_TRADE_TAG":                        55,
		"US_TX_PLUSPASS":                                56,
		"US_TX_TOLLTAG":                                 57,
		"US_TX_TXTAG":                                   58,
		"US_TX_XPRESS_CARD":                             59,
		"US_UT_ADAMS_AVE_PARKWAY_EXPRESSCARD":           60,
		"US_VA_EZPASSVA":                                61,
		"US_WA_BREEZEBY":                                17,
		"US_WA_GOOD_TO_GO":                              1,
		"US_WV_EZPASSWV":                                62,
		"US_WV_MEMORIAL_BRIDGE_TICKETS":                 63,
		"US_WV_NEWELL_TOLL_BRIDGE_TICKET":               64,
	}
)

func (x TollPass) Enum() *TollPass {
	p := new(TollPass)
	*p = x
	return p
}

func (x TollPass) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TollPass) Descriptor() protoreflect.EnumDescriptor {
	return file_google_maps_routing_v2_toll_passes_proto_enumTypes[0].Descriptor()
}

func (TollPass) Type() protoreflect.EnumType {
	return &file_google_maps_routing_v2_toll_passes_proto_enumTypes[0]
}

func (x TollPass) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TollPass.Descriptor instead.
func (TollPass) EnumDescriptor() ([]byte, []int) {
	return file_google_maps_routing_v2_toll_passes_proto_rawDescGZIP(), []int{0}
}

var File_google_maps_routing_v2_toll_passes_proto protoreflect.FileDescriptor

var file_google_maps_routing_v2_toll_passes_proto_rawDesc = []byte{
	0x0a, 0x28, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x72, 0x6f,
	0x75, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x32, 0x2f, 0x74, 0x6f, 0x6c, 0x6c, 0x5f, 0x70, 0x61,
	0x73, 0x73, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x32, 0x2a, 0x81, 0x11, 0x0a, 0x08, 0x54, 0x6f, 0x6c, 0x6c, 0x50, 0x61, 0x73, 0x73, 0x12,
	0x19, 0x0a, 0x15, 0x54, 0x4f, 0x4c, 0x4c, 0x5f, 0x50, 0x41, 0x53, 0x53, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x41, 0x55,
	0x5f, 0x45, 0x54, 0x4f, 0x4c, 0x4c, 0x5f, 0x54, 0x41, 0x47, 0x10, 0x52, 0x12, 0x0f, 0x0a, 0x0b,
	0x41, 0x55, 0x5f, 0x45, 0x57, 0x41, 0x59, 0x5f, 0x54, 0x41, 0x47, 0x10, 0x53, 0x12, 0x0c, 0x0a,
	0x08, 0x41, 0x55, 0x5f, 0x4c, 0x49, 0x4e, 0x4b, 0x54, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x41,
	0x52, 0x5f, 0x54, 0x45, 0x4c, 0x45, 0x50, 0x41, 0x53, 0x45, 0x10, 0x03, 0x12, 0x13, 0x0a, 0x0f,
	0x42, 0x52, 0x5f, 0x41, 0x55, 0x54, 0x4f, 0x5f, 0x45, 0x58, 0x50, 0x52, 0x45, 0x53, 0x4f, 0x10,
	0x51, 0x12, 0x10, 0x0a, 0x0c, 0x42, 0x52, 0x5f, 0x43, 0x4f, 0x4e, 0x45, 0x43, 0x54, 0x43, 0x41,
	0x52, 0x10, 0x07, 0x12, 0x10, 0x0a, 0x0c, 0x42, 0x52, 0x5f, 0x4d, 0x4f, 0x56, 0x45, 0x5f, 0x4d,
	0x41, 0x49, 0x53, 0x10, 0x08, 0x12, 0x13, 0x0a, 0x0f, 0x42, 0x52, 0x5f, 0x50, 0x41, 0x53, 0x53,
	0x41, 0x5f, 0x52, 0x41, 0x50, 0x49, 0x44, 0x4f, 0x10, 0x58, 0x12, 0x10, 0x0a, 0x0c, 0x42, 0x52,
	0x5f, 0x53, 0x45, 0x4d, 0x5f, 0x50, 0x41, 0x52, 0x41, 0x52, 0x10, 0x09, 0x12, 0x0c, 0x0a, 0x08,
	0x42, 0x52, 0x5f, 0x54, 0x41, 0x47, 0x47, 0x59, 0x10, 0x0a, 0x12, 0x0c, 0x0a, 0x08, 0x42, 0x52,
	0x5f, 0x56, 0x45, 0x4c, 0x4f, 0x45, 0x10, 0x0b, 0x12, 0x29, 0x0a, 0x25, 0x43, 0x41, 0x5f, 0x55,
	0x53, 0x5f, 0x41, 0x4b, 0x57, 0x41, 0x53, 0x41, 0x53, 0x4e, 0x45, 0x5f, 0x53, 0x45, 0x41, 0x57,
	0x41, 0x59, 0x5f, 0x43, 0x4f, 0x52, 0x50, 0x4f, 0x52, 0x41, 0x54, 0x45, 0x5f, 0x43, 0x41, 0x52,
	0x44, 0x10, 0x54, 0x12, 0x27, 0x0a, 0x23, 0x43, 0x41, 0x5f, 0x55, 0x53, 0x5f, 0x41, 0x4b, 0x57,
	0x41, 0x53, 0x41, 0x53, 0x4e, 0x45, 0x5f, 0x53, 0x45, 0x41, 0x57, 0x41, 0x59, 0x5f, 0x54, 0x52,
	0x41, 0x4e, 0x53, 0x49, 0x54, 0x5f, 0x43, 0x41, 0x52, 0x44, 0x10, 0x55, 0x12, 0x1e, 0x0a, 0x1a,
	0x43, 0x41, 0x5f, 0x55, 0x53, 0x5f, 0x42, 0x4c, 0x55, 0x45, 0x5f, 0x57, 0x41, 0x54, 0x45, 0x52,
	0x5f, 0x45, 0x44, 0x47, 0x45, 0x5f, 0x50, 0x41, 0x53, 0x53, 0x10, 0x12, 0x12, 0x13, 0x0a, 0x0f,
	0x43, 0x41, 0x5f, 0x55, 0x53, 0x5f, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x58, 0x49, 0x4f, 0x4e, 0x10,
	0x13, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x41, 0x5f, 0x55, 0x53, 0x5f, 0x4e, 0x45, 0x58, 0x55, 0x53,
	0x5f, 0x43, 0x41, 0x52, 0x44, 0x10, 0x14, 0x12, 0x0d, 0x0a, 0x09, 0x49, 0x44, 0x5f, 0x45, 0x5f,
	0x54, 0x4f, 0x4c, 0x4c, 0x10, 0x10, 0x12, 0x0d, 0x0a, 0x09, 0x49, 0x4e, 0x5f, 0x46, 0x41, 0x53,
	0x54, 0x41, 0x47, 0x10, 0x4e, 0x12, 0x1c, 0x0a, 0x18, 0x49, 0x4e, 0x5f, 0x4c, 0x4f, 0x43, 0x41,
	0x4c, 0x5f, 0x48, 0x50, 0x5f, 0x50, 0x4c, 0x41, 0x54, 0x45, 0x5f, 0x45, 0x58, 0x45, 0x4d, 0x50,
	0x54, 0x10, 0x4f, 0x12, 0x0b, 0x0a, 0x07, 0x4d, 0x58, 0x5f, 0x49, 0x41, 0x56, 0x45, 0x10, 0x5a,
	0x12, 0x0b, 0x0a, 0x07, 0x4d, 0x58, 0x5f, 0x50, 0x41, 0x53, 0x45, 0x10, 0x5b, 0x12, 0x10, 0x0a,
	0x0c, 0x4d, 0x58, 0x5f, 0x51, 0x55, 0x49, 0x43, 0x4b, 0x50, 0x41, 0x53, 0x53, 0x10, 0x5d, 0x12,
	0x22, 0x0a, 0x1e, 0x4d, 0x58, 0x5f, 0x53, 0x49, 0x53, 0x54, 0x45, 0x4d, 0x41, 0x5f, 0x54, 0x45,
	0x4c, 0x45, 0x50, 0x45, 0x41, 0x4a, 0x45, 0x5f, 0x43, 0x48, 0x49, 0x48, 0x55, 0x41, 0x48, 0x55,
	0x41, 0x10, 0x59, 0x12, 0x0f, 0x0a, 0x0b, 0x4d, 0x58, 0x5f, 0x54, 0x41, 0x47, 0x5f, 0x49, 0x41,
	0x56, 0x45, 0x10, 0x0c, 0x12, 0x12, 0x0a, 0x0e, 0x4d, 0x58, 0x5f, 0x54, 0x41, 0x47, 0x5f, 0x54,
	0x45, 0x4c, 0x45, 0x56, 0x49, 0x41, 0x10, 0x0d, 0x12, 0x0e, 0x0a, 0x0a, 0x4d, 0x58, 0x5f, 0x54,
	0x45, 0x4c, 0x45, 0x56, 0x49, 0x41, 0x10, 0x5c, 0x12, 0x0e, 0x0a, 0x0a, 0x4d, 0x58, 0x5f, 0x56,
	0x49, 0x41, 0x50, 0x41, 0x53, 0x53, 0x10, 0x0e, 0x12, 0x16, 0x0a, 0x12, 0x55, 0x53, 0x5f, 0x41,
	0x4c, 0x5f, 0x46, 0x52, 0x45, 0x45, 0x44, 0x4f, 0x4d, 0x5f, 0x50, 0x41, 0x53, 0x53, 0x10, 0x15,
	0x12, 0x32, 0x0a, 0x2e, 0x55, 0x53, 0x5f, 0x41, 0x4b, 0x5f, 0x41, 0x4e, 0x54, 0x4f, 0x4e, 0x5f,
	0x41, 0x4e, 0x44, 0x45, 0x52, 0x53, 0x4f, 0x4e, 0x5f, 0x54, 0x55, 0x4e, 0x4e, 0x45, 0x4c, 0x5f,
	0x42, 0x4f, 0x4f, 0x4b, 0x5f, 0x4f, 0x46, 0x5f, 0x31, 0x30, 0x5f, 0x54, 0x49, 0x43, 0x4b, 0x45,
	0x54, 0x53, 0x10, 0x16, 0x12, 0x11, 0x0a, 0x0d, 0x55, 0x53, 0x5f, 0x43, 0x41, 0x5f, 0x46, 0x41,
	0x53, 0x54, 0x52, 0x41, 0x4b, 0x10, 0x04, 0x12, 0x1d, 0x0a, 0x19, 0x55, 0x53, 0x5f, 0x43, 0x41,
	0x5f, 0x46, 0x41, 0x53, 0x54, 0x52, 0x41, 0x4b, 0x5f, 0x43, 0x41, 0x56, 0x5f, 0x53, 0x54, 0x49,
	0x43, 0x4b, 0x45, 0x52, 0x10, 0x56, 0x12, 0x15, 0x0a, 0x11, 0x55, 0x53, 0x5f, 0x43, 0x4f, 0x5f,
	0x45, 0x58, 0x50, 0x52, 0x45, 0x53, 0x53, 0x54, 0x4f, 0x4c, 0x4c, 0x10, 0x17, 0x12, 0x11, 0x0a,
	0x0d, 0x55, 0x53, 0x5f, 0x43, 0x4f, 0x5f, 0x47, 0x4f, 0x5f, 0x50, 0x41, 0x53, 0x53, 0x10, 0x18,
	0x12, 0x12, 0x0a, 0x0e, 0x55, 0x53, 0x5f, 0x44, 0x45, 0x5f, 0x45, 0x5a, 0x50, 0x41, 0x53, 0x53,
	0x44, 0x45, 0x10, 0x19, 0x12, 0x24, 0x0a, 0x20, 0x55, 0x53, 0x5f, 0x46, 0x4c, 0x5f, 0x42, 0x4f,
	0x42, 0x5f, 0x53, 0x49, 0x4b, 0x45, 0x53, 0x5f, 0x54, 0x4f, 0x4c, 0x4c, 0x5f, 0x42, 0x52, 0x49,
	0x44, 0x47, 0x45, 0x5f, 0x50, 0x41, 0x53, 0x53, 0x10, 0x41, 0x12, 0x3a, 0x0a, 0x36, 0x55, 0x53,
	0x5f, 0x46, 0x4c, 0x5f, 0x44, 0x55, 0x4e, 0x45, 0x53, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x55, 0x4e,
	0x49, 0x54, 0x59, 0x5f, 0x44, 0x45, 0x56, 0x45, 0x4c, 0x4f, 0x50, 0x4d, 0x45, 0x4e, 0x54, 0x5f,
	0x44, 0x49, 0x53, 0x54, 0x52, 0x49, 0x43, 0x54, 0x5f, 0x45, 0x58, 0x50, 0x52, 0x45, 0x53, 0x53,
	0x43, 0x41, 0x52, 0x44, 0x10, 0x42, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x53, 0x5f, 0x46, 0x4c, 0x5f,
	0x45, 0x50, 0x41, 0x53, 0x53, 0x10, 0x43, 0x12, 0x18, 0x0a, 0x14, 0x55, 0x53, 0x5f, 0x46, 0x4c,
	0x5f, 0x47, 0x49, 0x42, 0x41, 0x5f, 0x54, 0x4f, 0x4c, 0x4c, 0x5f, 0x50, 0x41, 0x53, 0x53, 0x10,
	0x44, 0x12, 0x10, 0x0a, 0x0c, 0x55, 0x53, 0x5f, 0x46, 0x4c, 0x5f, 0x4c, 0x45, 0x45, 0x57, 0x41,
	0x59, 0x10, 0x45, 0x12, 0x11, 0x0a, 0x0d, 0x55, 0x53, 0x5f, 0x46, 0x4c, 0x5f, 0x53, 0x55, 0x4e,
	0x50, 0x41, 0x53, 0x53, 0x10, 0x46, 0x12, 0x15, 0x0a, 0x11, 0x55, 0x53, 0x5f, 0x46, 0x4c, 0x5f,
	0x53, 0x55, 0x4e, 0x50, 0x41, 0x53, 0x53, 0x5f, 0x50, 0x52, 0x4f, 0x10, 0x47, 0x12, 0x12, 0x0a,
	0x0e, 0x55, 0x53, 0x5f, 0x49, 0x4c, 0x5f, 0x45, 0x5a, 0x50, 0x41, 0x53, 0x53, 0x49, 0x4c, 0x10,
	0x49, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x53, 0x5f, 0x49, 0x4c, 0x5f, 0x49, 0x50, 0x41, 0x53, 0x53,
	0x10, 0x48, 0x12, 0x12, 0x0a, 0x0e, 0x55, 0x53, 0x5f, 0x49, 0x4e, 0x5f, 0x45, 0x5a, 0x50, 0x41,
	0x53, 0x53, 0x49, 0x4e, 0x10, 0x1a, 0x12, 0x1a, 0x0a, 0x16, 0x55, 0x53, 0x5f, 0x4b, 0x53, 0x5f,
	0x42, 0x45, 0x53, 0x54, 0x50, 0x41, 0x53, 0x53, 0x5f, 0x48, 0x4f, 0x52, 0x49, 0x5a, 0x4f, 0x4e,
	0x10, 0x1b, 0x12, 0x0e, 0x0a, 0x0a, 0x55, 0x53, 0x5f, 0x4b, 0x53, 0x5f, 0x4b, 0x54, 0x41, 0x47,
	0x10, 0x1c, 0x12, 0x16, 0x0a, 0x12, 0x55, 0x53, 0x5f, 0x4b, 0x53, 0x5f, 0x4e, 0x41, 0x54, 0x49,
	0x4f, 0x4e, 0x41, 0x4c, 0x50, 0x41, 0x53, 0x53, 0x10, 0x1d, 0x12, 0x1b, 0x0a, 0x17, 0x55, 0x53,
	0x5f, 0x4b, 0x53, 0x5f, 0x50, 0x52, 0x45, 0x50, 0x41, 0x53, 0x53, 0x5f, 0x45, 0x4c, 0x49, 0x54,
	0x45, 0x50, 0x41, 0x53, 0x53, 0x10, 0x1e, 0x12, 0x13, 0x0a, 0x0f, 0x55, 0x53, 0x5f, 0x4b, 0x59,
	0x5f, 0x52, 0x49, 0x56, 0x45, 0x52, 0x4c, 0x49, 0x4e, 0x4b, 0x10, 0x1f, 0x12, 0x13, 0x0a, 0x0f,
	0x55, 0x53, 0x5f, 0x4c, 0x41, 0x5f, 0x47, 0x45, 0x41, 0x55, 0x58, 0x50, 0x41, 0x53, 0x53, 0x10,
	0x20, 0x12, 0x12, 0x0a, 0x0e, 0x55, 0x53, 0x5f, 0x4c, 0x41, 0x5f, 0x54, 0x4f, 0x4c, 0x4c, 0x5f,
	0x54, 0x41, 0x47, 0x10, 0x21, 0x12, 0x12, 0x0a, 0x0e, 0x55, 0x53, 0x5f, 0x4d, 0x41, 0x5f, 0x45,
	0x5a, 0x50, 0x41, 0x53, 0x53, 0x4d, 0x41, 0x10, 0x06, 0x12, 0x12, 0x0a, 0x0e, 0x55, 0x53, 0x5f,
	0x4d, 0x44, 0x5f, 0x45, 0x5a, 0x50, 0x41, 0x53, 0x53, 0x4d, 0x44, 0x10, 0x22, 0x12, 0x12, 0x0a,
	0x0e, 0x55, 0x53, 0x5f, 0x4d, 0x45, 0x5f, 0x45, 0x5a, 0x50, 0x41, 0x53, 0x53, 0x4d, 0x45, 0x10,
	0x23, 0x12, 0x31, 0x0a, 0x2d, 0x55, 0x53, 0x5f, 0x4d, 0x49, 0x5f, 0x41, 0x4d, 0x42, 0x41, 0x53,
	0x53, 0x41, 0x44, 0x4f, 0x52, 0x5f, 0x42, 0x52, 0x49, 0x44, 0x47, 0x45, 0x5f, 0x50, 0x52, 0x45,
	0x4d, 0x49, 0x45, 0x52, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x55, 0x54, 0x45, 0x52, 0x5f, 0x43, 0x41,
	0x52, 0x44, 0x10, 0x24, 0x12, 0x29, 0x0a, 0x25, 0x55, 0x53, 0x5f, 0x4d, 0x49, 0x5f, 0x47, 0x52,
	0x4f, 0x53, 0x53, 0x45, 0x5f, 0x49, 0x4c, 0x45, 0x5f, 0x54, 0x4f, 0x4c, 0x4c, 0x5f, 0x42, 0x52,
	0x49, 0x44, 0x47, 0x45, 0x5f, 0x50, 0x41, 0x53, 0x53, 0x5f, 0x54, 0x41, 0x47, 0x10, 0x25, 0x12,
	0x16, 0x0a, 0x12, 0x55, 0x53, 0x5f, 0x4d, 0x49, 0x5f, 0x49, 0x51, 0x5f, 0x50, 0x52, 0x4f, 0x58,
	0x5f, 0x43, 0x41, 0x52, 0x44, 0x10, 0x26, 0x12, 0x22, 0x0a, 0x1e, 0x55, 0x53, 0x5f, 0x4d, 0x49,
	0x5f, 0x4d, 0x41, 0x43, 0x4b, 0x49, 0x4e, 0x41, 0x43, 0x5f, 0x42, 0x52, 0x49, 0x44, 0x47, 0x45,
	0x5f, 0x4d, 0x41, 0x43, 0x5f, 0x50, 0x41, 0x53, 0x53, 0x10, 0x27, 0x12, 0x17, 0x0a, 0x13, 0x55,
	0x53, 0x5f, 0x4d, 0x49, 0x5f, 0x4e, 0x45, 0x58, 0x50, 0x52, 0x45, 0x53, 0x53, 0x5f, 0x54, 0x4f,
	0x4c, 0x4c, 0x10, 0x28, 0x12, 0x12, 0x0a, 0x0e, 0x55, 0x53, 0x5f, 0x4d, 0x4e, 0x5f, 0x45, 0x5a,
	0x50, 0x41, 0x53, 0x53, 0x4d, 0x4e, 0x10, 0x29, 0x12, 0x12, 0x0a, 0x0e, 0x55, 0x53, 0x5f, 0x4e,
	0x43, 0x5f, 0x45, 0x5a, 0x50, 0x41, 0x53, 0x53, 0x4e, 0x43, 0x10, 0x2a, 0x12, 0x14, 0x0a, 0x10,
	0x55, 0x53, 0x5f, 0x4e, 0x43, 0x5f, 0x50, 0x45, 0x41, 0x43, 0x48, 0x5f, 0x50, 0x41, 0x53, 0x53,
	0x10, 0x57, 0x12, 0x14, 0x0a, 0x10, 0x55, 0x53, 0x5f, 0x4e, 0x43, 0x5f, 0x51, 0x55, 0x49, 0x43,
	0x4b, 0x5f, 0x50, 0x41, 0x53, 0x53, 0x10, 0x2b, 0x12, 0x12, 0x0a, 0x0e, 0x55, 0x53, 0x5f, 0x4e,
	0x48, 0x5f, 0x45, 0x5a, 0x50, 0x41, 0x53, 0x53, 0x4e, 0x48, 0x10, 0x50, 0x12, 0x20, 0x0a, 0x1c,
	0x55, 0x53, 0x5f, 0x4e, 0x4a, 0x5f, 0x44, 0x4f, 0x57, 0x4e, 0x42, 0x45, 0x41, 0x43, 0x48, 0x5f,
	0x45, 0x58, 0x50, 0x52, 0x45, 0x53, 0x53, 0x5f, 0x50, 0x41, 0x53, 0x53, 0x10, 0x4b, 0x12, 0x12,
	0x0a, 0x0e, 0x55, 0x53, 0x5f, 0x4e, 0x4a, 0x5f, 0x45, 0x5a, 0x50, 0x41, 0x53, 0x53, 0x4e, 0x4a,
	0x10, 0x4a, 0x12, 0x15, 0x0a, 0x11, 0x55, 0x53, 0x5f, 0x4e, 0x59, 0x5f, 0x45, 0x58, 0x50, 0x52,
	0x45, 0x53, 0x53, 0x50, 0x41, 0x53, 0x53, 0x10, 0x4c, 0x12, 0x12, 0x0a, 0x0e, 0x55, 0x53, 0x5f,
	0x4e, 0x59, 0x5f, 0x45, 0x5a, 0x50, 0x41, 0x53, 0x53, 0x4e, 0x59, 0x10, 0x4d, 0x12, 0x12, 0x0a,
	0x0e, 0x55, 0x53, 0x5f, 0x4f, 0x48, 0x5f, 0x45, 0x5a, 0x50, 0x41, 0x53, 0x53, 0x4f, 0x48, 0x10,
	0x2c, 0x12, 0x12, 0x0a, 0x0e, 0x55, 0x53, 0x5f, 0x50, 0x41, 0x5f, 0x45, 0x5a, 0x50, 0x41, 0x53,
	0x53, 0x50, 0x41, 0x10, 0x2d, 0x12, 0x12, 0x0a, 0x0e, 0x55, 0x53, 0x5f, 0x52, 0x49, 0x5f, 0x45,
	0x5a, 0x50, 0x41, 0x53, 0x53, 0x52, 0x49, 0x10, 0x2e, 0x12, 0x11, 0x0a, 0x0d, 0x55, 0x53, 0x5f,
	0x53, 0x43, 0x5f, 0x50, 0x41, 0x4c, 0x50, 0x41, 0x53, 0x53, 0x10, 0x2f, 0x12, 0x12, 0x0a, 0x0e,
	0x55, 0x53, 0x5f, 0x54, 0x58, 0x5f, 0x42, 0x41, 0x4e, 0x43, 0x50, 0x41, 0x53, 0x53, 0x10, 0x30,
	0x12, 0x16, 0x0a, 0x12, 0x55, 0x53, 0x5f, 0x54, 0x58, 0x5f, 0x44, 0x45, 0x4c, 0x5f, 0x52, 0x49,
	0x4f, 0x5f, 0x50, 0x41, 0x53, 0x53, 0x10, 0x31, 0x12, 0x14, 0x0a, 0x10, 0x55, 0x53, 0x5f, 0x54,
	0x58, 0x5f, 0x45, 0x46, 0x41, 0x53, 0x54, 0x5f, 0x50, 0x41, 0x53, 0x53, 0x10, 0x32, 0x12, 0x21,
	0x0a, 0x1d, 0x55, 0x53, 0x5f, 0x54, 0x58, 0x5f, 0x45, 0x41, 0x47, 0x4c, 0x45, 0x5f, 0x50, 0x41,
	0x53, 0x53, 0x5f, 0x45, 0x58, 0x50, 0x52, 0x45, 0x53, 0x53, 0x5f, 0x43, 0x41, 0x52, 0x44, 0x10,
	0x33, 0x12, 0x10, 0x0a, 0x0c, 0x55, 0x53, 0x5f, 0x54, 0x58, 0x5f, 0x45, 0x50, 0x54, 0x4f, 0x4c,
	0x4c, 0x10, 0x34, 0x12, 0x12, 0x0a, 0x0e, 0x55, 0x53, 0x5f, 0x54, 0x58, 0x5f, 0x45, 0x5a, 0x5f,
	0x43, 0x52, 0x4f, 0x53, 0x53, 0x10, 0x35, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x53, 0x5f, 0x54, 0x58,
	0x5f, 0x45, 0x5a, 0x54, 0x41, 0x47, 0x10, 0x36, 0x12, 0x1a, 0x0a, 0x16, 0x55, 0x53, 0x5f, 0x54,
	0x58, 0x5f, 0x4c, 0x41, 0x52, 0x45, 0x44, 0x4f, 0x5f, 0x54, 0x52, 0x41, 0x44, 0x45, 0x5f, 0x54,
	0x41, 0x47, 0x10, 0x37, 0x12, 0x12, 0x0a, 0x0e, 0x55, 0x53, 0x5f, 0x54, 0x58, 0x5f, 0x50, 0x4c,
	0x55, 0x53, 0x50, 0x41, 0x53, 0x53, 0x10, 0x38, 0x12, 0x11, 0x0a, 0x0d, 0x55, 0x53, 0x5f, 0x54,
	0x58, 0x5f, 0x54, 0x4f, 0x4c, 0x4c, 0x54, 0x41, 0x47, 0x10, 0x39, 0x12, 0x0f, 0x0a, 0x0b, 0x55,
	0x53, 0x5f, 0x54, 0x58, 0x5f, 0x54, 0x58, 0x54, 0x41, 0x47, 0x10, 0x3a, 0x12, 0x15, 0x0a, 0x11,
	0x55, 0x53, 0x5f, 0x54, 0x58, 0x5f, 0x58, 0x50, 0x52, 0x45, 0x53, 0x53, 0x5f, 0x43, 0x41, 0x52,
	0x44, 0x10, 0x3b, 0x12, 0x27, 0x0a, 0x23, 0x55, 0x53, 0x5f, 0x55, 0x54, 0x5f, 0x41, 0x44, 0x41,
	0x4d, 0x53, 0x5f, 0x41, 0x56, 0x45, 0x5f, 0x50, 0x41, 0x52, 0x4b, 0x57, 0x41, 0x59, 0x5f, 0x45,
	0x58, 0x50, 0x52, 0x45, 0x53, 0x53, 0x43, 0x41, 0x52, 0x44, 0x10, 0x3c, 0x12, 0x12, 0x0a, 0x0e,
	0x55, 0x53, 0x5f, 0x56, 0x41, 0x5f, 0x45, 0x5a, 0x50, 0x41, 0x53, 0x53, 0x56, 0x41, 0x10, 0x3d,
	0x12, 0x12, 0x0a, 0x0e, 0x55, 0x53, 0x5f, 0x57, 0x41, 0x5f, 0x42, 0x52, 0x45, 0x45, 0x5a, 0x45,
	0x42, 0x59, 0x10, 0x11, 0x12, 0x14, 0x0a, 0x10, 0x55, 0x53, 0x5f, 0x57, 0x41, 0x5f, 0x47, 0x4f,
	0x4f, 0x44, 0x5f, 0x54, 0x4f, 0x5f, 0x47, 0x4f, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x55, 0x53,
	0x5f, 0x57, 0x56, 0x5f, 0x45, 0x5a, 0x50, 0x41, 0x53, 0x53, 0x57, 0x56, 0x10, 0x3e, 0x12, 0x21,
	0x0a, 0x1d, 0x55, 0x53, 0x5f, 0x57, 0x56, 0x5f, 0x4d, 0x45, 0x4d, 0x4f, 0x52, 0x49, 0x41, 0x4c,
	0x5f, 0x42, 0x52, 0x49, 0x44, 0x47, 0x45, 0x5f, 0x54, 0x49, 0x43, 0x4b, 0x45, 0x54, 0x53, 0x10,
	0x3f, 0x12, 0x23, 0x0a, 0x1f, 0x55, 0x53, 0x5f, 0x57, 0x56, 0x5f, 0x4e, 0x45, 0x57, 0x45, 0x4c,
	0x4c, 0x5f, 0x54, 0x4f, 0x4c, 0x4c, 0x5f, 0x42, 0x52, 0x49, 0x44, 0x47, 0x45, 0x5f, 0x54, 0x49,
	0x43, 0x4b, 0x45, 0x54, 0x10, 0x40, 0x42, 0xc7, 0x01, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69,
	0x6e, 0x67, 0x2e, 0x76, 0x32, 0x42, 0x0f, 0x54, 0x6f, 0x6c, 0x6c, 0x50, 0x61, 0x73, 0x73, 0x65,
	0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x6d, 0x61, 0x70, 0x73, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x32, 0x3b,
	0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0xf8, 0x01, 0x01, 0xa2, 0x02, 0x05, 0x47, 0x4d, 0x52,
	0x56, 0x32, 0xaa, 0x02, 0x16, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x4d, 0x61, 0x70, 0x73,
	0x2e, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x32, 0xca, 0x02, 0x16, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x4d, 0x61, 0x70, 0x73, 0x5c, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e,
	0x67, 0x5c, 0x56, 0x32, 0xea, 0x02, 0x19, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x4d,
	0x61, 0x70, 0x73, 0x3a, 0x3a, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x3a, 0x3a, 0x56, 0x32,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_maps_routing_v2_toll_passes_proto_rawDescOnce sync.Once
	file_google_maps_routing_v2_toll_passes_proto_rawDescData = file_google_maps_routing_v2_toll_passes_proto_rawDesc
)

func file_google_maps_routing_v2_toll_passes_proto_rawDescGZIP() []byte {
	file_google_maps_routing_v2_toll_passes_proto_rawDescOnce.Do(func() {
		file_google_maps_routing_v2_toll_passes_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_maps_routing_v2_toll_passes_proto_rawDescData)
	})
	return file_google_maps_routing_v2_toll_passes_proto_rawDescData
}

var file_google_maps_routing_v2_toll_passes_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_maps_routing_v2_toll_passes_proto_goTypes = []interface{}{
	(TollPass)(0), // 0: google.maps.routing.v2.TollPass
}
var file_google_maps_routing_v2_toll_passes_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_maps_routing_v2_toll_passes_proto_init() }
func file_google_maps_routing_v2_toll_passes_proto_init() {
	if File_google_maps_routing_v2_toll_passes_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_maps_routing_v2_toll_passes_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_maps_routing_v2_toll_passes_proto_goTypes,
		DependencyIndexes: file_google_maps_routing_v2_toll_passes_proto_depIdxs,
		EnumInfos:         file_google_maps_routing_v2_toll_passes_proto_enumTypes,
	}.Build()
	File_google_maps_routing_v2_toll_passes_proto = out.File
	file_google_maps_routing_v2_toll_passes_proto_rawDesc = nil
	file_google_maps_routing_v2_toll_passes_proto_goTypes = nil
	file_google_maps_routing_v2_toll_passes_proto_depIdxs = nil
}
