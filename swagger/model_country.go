/*
 * Square
 *
 * Use Square APIs to manage and run business including payment, customer, product, inventory, and employee management.
 *
 * API version: 2.0
 * Contact: developers@squareup.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Country : Indicates the country associated with another entity, such as a business. Values are in [ISO 3166-1-alpha-2 format](http://www.iso.org/iso/home/standards/country_codes.htm).
type Country string

// List of Country
const (
	ZZ_Country Country = "ZZ"
	AD_Country Country = "AD"
	AE_Country Country = "AE"
	AF_Country Country = "AF"
	AG_Country Country = "AG"
	AI_Country Country = "AI"
	AL_Country Country = "AL"
	AM_Country Country = "AM"
	AO_Country Country = "AO"
	AQ_Country Country = "AQ"
	AR_Country Country = "AR"
	AS_Country Country = "AS"
	AT_Country Country = "AT"
	AU_Country Country = "AU"
	AW_Country Country = "AW"
	AX_Country Country = "AX"
	AZ_Country Country = "AZ"
	BA_Country Country = "BA"
	BB_Country Country = "BB"
	BD_Country Country = "BD"
	BE_Country Country = "BE"
	BF_Country Country = "BF"
	BG_Country Country = "BG"
	BH_Country Country = "BH"
	BI_Country Country = "BI"
	BJ_Country Country = "BJ"
	BL_Country Country = "BL"
	BM_Country Country = "BM"
	BN_Country Country = "BN"
	BO_Country Country = "BO"
	BQ_Country Country = "BQ"
	BR_Country Country = "BR"
	BS_Country Country = "BS"
	BT_Country Country = "BT"
	BV_Country Country = "BV"
	BW_Country Country = "BW"
	BY_Country Country = "BY"
	BZ_Country Country = "BZ"
	CA_Country Country = "CA"
	CC_Country Country = "CC"
	CD_Country Country = "CD"
	CF_Country Country = "CF"
	CG_Country Country = "CG"
	CH_Country Country = "CH"
	CI_Country Country = "CI"
	CK_Country Country = "CK"
	CL_Country Country = "CL"
	CM_Country Country = "CM"
	CN_Country Country = "CN"
	CO_Country Country = "CO"
	CR_Country Country = "CR"
	CU_Country Country = "CU"
	CV_Country Country = "CV"
	CW_Country Country = "CW"
	CX_Country Country = "CX"
	CY_Country Country = "CY"
	CZ_Country Country = "CZ"
	DE_Country Country = "DE"
	DJ_Country Country = "DJ"
	DK_Country Country = "DK"
	DM_Country Country = "DM"
	DO_Country Country = "DO"
	DZ_Country Country = "DZ"
	EC_Country Country = "EC"
	EE_Country Country = "EE"
	EG_Country Country = "EG"
	EH_Country Country = "EH"
	ER_Country Country = "ER"
	ES_Country Country = "ES"
	ET_Country Country = "ET"
	FI_Country Country = "FI"
	FJ_Country Country = "FJ"
	FK_Country Country = "FK"
	FM_Country Country = "FM"
	FO_Country Country = "FO"
	FR_Country Country = "FR"
	GA_Country Country = "GA"
	GB_Country Country = "GB"
	GD_Country Country = "GD"
	GE_Country Country = "GE"
	GF_Country Country = "GF"
	GG_Country Country = "GG"
	GH_Country Country = "GH"
	GI_Country Country = "GI"
	GL_Country Country = "GL"
	GM_Country Country = "GM"
	GN_Country Country = "GN"
	GP_Country Country = "GP"
	GQ_Country Country = "GQ"
	GR_Country Country = "GR"
	GS_Country Country = "GS"
	GT_Country Country = "GT"
	GU_Country Country = "GU"
	GW_Country Country = "GW"
	GY_Country Country = "GY"
	HK_Country Country = "HK"
	HM_Country Country = "HM"
	HN_Country Country = "HN"
	HR_Country Country = "HR"
	HT_Country Country = "HT"
	HU_Country Country = "HU"
	ID_Country Country = "ID"
	IE_Country Country = "IE"
	IL_Country Country = "IL"
	IM_Country Country = "IM"
	IN_Country Country = "IN"
	IO_Country Country = "IO"
	IQ_Country Country = "IQ"
	IR_Country Country = "IR"
	IS_Country Country = "IS"
	IT_Country Country = "IT"
	JE_Country Country = "JE"
	JM_Country Country = "JM"
	JO_Country Country = "JO"
	JP_Country Country = "JP"
	KE_Country Country = "KE"
	KG_Country Country = "KG"
	KH_Country Country = "KH"
	KI_Country Country = "KI"
	KM_Country Country = "KM"
	KN_Country Country = "KN"
	KP_Country Country = "KP"
	KR_Country Country = "KR"
	KW_Country Country = "KW"
	KY_Country Country = "KY"
	KZ_Country Country = "KZ"
	LA_Country Country = "LA"
	LB_Country Country = "LB"
	LC_Country Country = "LC"
	LI_Country Country = "LI"
	LK_Country Country = "LK"
	LR_Country Country = "LR"
	LS_Country Country = "LS"
	LT_Country Country = "LT"
	LU_Country Country = "LU"
	LV_Country Country = "LV"
	LY_Country Country = "LY"
	MA_Country Country = "MA"
	MC_Country Country = "MC"
	MD_Country Country = "MD"
	ME_Country Country = "ME"
	MF_Country Country = "MF"
	MG_Country Country = "MG"
	MH_Country Country = "MH"
	MK_Country Country = "MK"
	ML_Country Country = "ML"
	MM_Country Country = "MM"
	MN_Country Country = "MN"
	MO_Country Country = "MO"
	MP_Country Country = "MP"
	MQ_Country Country = "MQ"
	MR_Country Country = "MR"
	MS_Country Country = "MS"
	MT_Country Country = "MT"
	MU_Country Country = "MU"
	MV_Country Country = "MV"
	MW_Country Country = "MW"
	MX_Country Country = "MX"
	MY_Country Country = "MY"
	MZ_Country Country = "MZ"
	NA_Country Country = "NA"
	NC_Country Country = "NC"
	NE_Country Country = "NE"
	NF_Country Country = "NF"
	NG_Country Country = "NG"
	NI_Country Country = "NI"
	NL_Country Country = "NL"
	NO_Country Country = "NO"
	NP_Country Country = "NP"
	NR_Country Country = "NR"
	NU_Country Country = "NU"
	NZ_Country Country = "NZ"
	OM_Country Country = "OM"
	PA_Country Country = "PA"
	PE_Country Country = "PE"
	PF_Country Country = "PF"
	PG_Country Country = "PG"
	PH_Country Country = "PH"
	PK_Country Country = "PK"
	PL_Country Country = "PL"
	PM_Country Country = "PM"
	PN_Country Country = "PN"
	PR_Country Country = "PR"
	PS_Country Country = "PS"
	PT_Country Country = "PT"
	PW_Country Country = "PW"
	PY_Country Country = "PY"
	QA_Country Country = "QA"
	RE_Country Country = "RE"
	RO_Country Country = "RO"
	RS_Country Country = "RS"
	RU_Country Country = "RU"
	RW_Country Country = "RW"
	SA_Country Country = "SA"
	SB_Country Country = "SB"
	SC_Country Country = "SC"
	SD_Country Country = "SD"
	SE_Country Country = "SE"
	SG_Country Country = "SG"
	SH_Country Country = "SH"
	SI_Country Country = "SI"
	SJ_Country Country = "SJ"
	SK_Country Country = "SK"
	SL_Country Country = "SL"
	SM_Country Country = "SM"
	SN_Country Country = "SN"
	SO_Country Country = "SO"
	SR_Country Country = "SR"
	SS_Country Country = "SS"
	ST_Country Country = "ST"
	SV_Country Country = "SV"
	SX_Country Country = "SX"
	SY_Country Country = "SY"
	SZ_Country Country = "SZ"
	TC_Country Country = "TC"
	TD_Country Country = "TD"
	TF_Country Country = "TF"
	TG_Country Country = "TG"
	TH_Country Country = "TH"
	TJ_Country Country = "TJ"
	TK_Country Country = "TK"
	TL_Country Country = "TL"
	TM_Country Country = "TM"
	TN_Country Country = "TN"
	TO_Country Country = "TO"
	TR_Country Country = "TR"
	TT_Country Country = "TT"
	TV_Country Country = "TV"
	TW_Country Country = "TW"
	TZ_Country Country = "TZ"
	UA_Country Country = "UA"
	UG_Country Country = "UG"
	UM_Country Country = "UM"
	US_Country Country = "US"
	UY_Country Country = "UY"
	UZ_Country Country = "UZ"
	VA_Country Country = "VA"
	VC_Country Country = "VC"
	VE_Country Country = "VE"
	VG_Country Country = "VG"
	VI_Country Country = "VI"
	VN_Country Country = "VN"
	VU_Country Country = "VU"
	WF_Country Country = "WF"
	WS_Country Country = "WS"
	YE_Country Country = "YE"
	YT_Country Country = "YT"
	ZA_Country Country = "ZA"
	ZM_Country Country = "ZM"
	ZW_Country Country = "ZW"
)
