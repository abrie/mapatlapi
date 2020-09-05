package record

import ()

type Submitter interface {
	Submit(r *Request) (*Result, error)
}

type Request struct {
	Ref_ID int
}

type Record struct {
	OBJECTID                string `json:"OBJECTID,omitempty"`
	FEATUREID               string `json:"FEATUREID,omitempty"`
	PARCELPIN               string `json:"PARCELPIN,omitempty"`
	STREETID                string `json:"STREETID,omitempty"`
	PREADDRNUM              string `json:"PREADDRNUM,omitempty"`
	STNUMBER                string `json:"STNUMBER,omitempty"`
	PREMOD                  string `json:"PREMOD,omitempty"`
	PREDIR                  string `json:"PREDIR,omitempty"`
	POSTMOD                 string `json:"POSTMOD,omitempty"`
	PRETYPE                 string `json:"PRETYPE,omitempty"`
	POSTDIR                 string `json:"POSTDIR,omitempty"`
	LABEL                   string `json:"LABEL,omitempty"`
	UNITLOW                 string `json:"UNITLOW,omitempty"`
	UNITHIGH                string `json:"UNITHIGH,omitempty"`
	UNITACTUAL              string `json:"UNITACTUAL,omitempty"`
	CITY                    string `json:"CITY,omitempty"`
	ZIP                     string `json:"ZIP,omitempty"`
	ZIP4                    string `json:"ZIP4,omitempty"`
	PTTYPE                  string `json:"PTTYPE,omitempty"`
	PTTYPESUB               string `json:"PTTYPESUB,omitempty"`
	SUBDIVISION             string `json:"SUBDIVISION,omitempty"`
	STATUS                  string `json:"STATUS,omitempty"`
	COMMONNAME              string `json:"COMMONNAME,omitempty"`
	SOURCE                  string `json:"SOURCE,omitempty"`
	COMMENT                 string `json:"COMMENT_,omitempty"`
	CREATEDATE              string `json:"CREATEDATE,omitempty"`
	STRNUMBERHI             string `json:"STRNUMBERHI,omitempty"`
	W_METER_NUM             string `json:"W_METER_NUM,omitempty"`
	W_PREMISE_ID            string `json:"W_PREMISE_ID,omitempty"`
	W_METER_TYPE            string `json:"W_METER_TYPE,omitempty"`
	W_METER_IN_CITY         string `json:"W_METER_IN_CITY,omitempty"`
	W_GT_NAME               string `json:"W_GT_NAME,omitempty"`
	W_GT_CITY_INSP          string `json:"W_GT_CITY_INSP,omitempty"`
	W_SP_TANK_ADDR          string `json:"W_SP_TANK_ADDR,omitempty"`
	W_SP_TO_CITYSEWER       string `json:"W_SP_TO_CITYSEWER,omitempty"`
	W_SP_AVAIL_TO_CONN      string `json:"W_SP_AVAIL_TO_CONN,omitempty"`
	W_MH_IN_CITY            string `json:"W_MH_IN_CITY,omitempty"`
	W_SWB_IN_CITY           string `json:"W_SWB_IN_CITY,omitempty"`
	W_HYDRANT_IN_CITY       string `json:"W_HYDRANT_IN_CITY,omitempty"`
	BLDG_ZONE               string `json:"BLDG_ZONE,omitempty"`
	BLDG_INSP_DISPLINE      string `json:"BLDG_INSP_DISPLINE,omitempty"`
	BLDG_DCA                string `json:"BLDG_DCA,omitempty"`
	BLDG_INSPR_NAME         string `json:"BLDG_INSPR_NAME,omitempty"`
	ELEM_SCH_ZONE           string `json:"ELEM_SCH_ZONE,omitempty"`
	HIGH_SCH_ZONE           string `json:"HIGH_SCH_ZONE,omitempty"`
	MID_SCH_ZONE            string `json:"MID_SCH_ZONE,omitempty"`
	LAND_LOT_NUM            string `json:"LAND_LOT_NUM,omitempty"`
	TAXPIN                  string `json:"TAXPIN,omitempty"`
	PRSA_NAME               string `json:"PRSA_NAME,omitempty"`
	NPU_NAME                string `json:"NPU_NAME,omitempty"`
	CDBG_NAME               string `json:"CDBGNAME,omitempty"`
	DIST_ZONING             string `json:"DIST_ZONING,omitempty"`
	OL_DIST_ZONING          string `json:"OL_DIST_ZONING,omitempty"`
	ADDRESS_IN_CITY         string `json:"ADDRESS_IN_CITY,omitempty"`
	COUNCIL_DIST            string `json:"COUNCIL_DIST,omitempty"`
	COUNCIL_MEMBER          string `json:"COUNCIL_MEMBER,omitempty"`
	VOTE_PRCTS              string `json:"VOTE_PRCTS,omitempty"`
	POLICE_ZONE             string `json:"POLICE_ZONE,omitempty"`
	POLICE_ST_COUNTY        string `json:"POLICE_ST_COUNTY,omitempty"`
	FIRE_ST_NAME            string `json:"FIRE_ST_NAME,omitempty"`
	FIRE_ST_ADDR            string `json:"FIRE_ST_ADDR,omitempty"`
	FIRESTTERR              string `json:"FIRE_ST_TERR,omitempty"`
	FIRE_BATTALION          string `json:"FIRE_BATTALION,omitempty"`
	PARK_NAME               string `json:"PARK_NAME,omitempty"`
	PARK_ADDR               string `json:"PARK_ADDR,omitempty"`
	FULL_STREET_ADDRESS     string `json:"FULL_STREET_ADDRESS,omitempty"`
	STREET_NUMBER           string `json:"STREET_NUMBER,omitempty"`
	STREET_PREFIX           string `json:"STREET_PREFIX,omitempty"`
	STREET_NAME             string `json:"STREET_NAME,omitempty"`
	STREET_TYPE             string `json:"STREET_TYPE,omitempty"`
	STREET_SUFFIX           string `json:"STREET_SUFFIX,omitempty"`
	UNIT_NUMBER             string `json:"UNIT_NUMBER,omitempty"`
	STATE                   string `json:"STATE,omitempty"`
	ZIPCODE                 string `json:"ZIPCODE,omitempty"`
	LATITUDE                string `json:"LATITUDE,omitempty"`
	LONGITUDE               string `json:"LONGITUDE,omitempty"`
	E311_UNIQUE_ID          string `json:"E311_UNIQUE_ID,omitempty"`
	TAX_PIN_ORIGINAL        string `json:"TAXPIN_ORIGINAL,omitempty"`
	NEAR_PARK_DISTANCE      string `json:"NEAR_PARK_DISTANCE,omitempty"`
	NEAR_POLICE_ST_FID      string `json:"NEAR_POLICE_ST_FID,omitempty"`
	NEAR_POLICE_ST_NUMBER   string `json:"NEAR_POLICE_ST_NUMBER,omitempty"`
	NEAR_POLICE_ST_ADDR     string `json:"NEAR_POLICE_ST_ADDR,omitempty"`
	NEAR_HYDRANT_ID         string `json:"NEAR_HYDRANT_ID,omitempty"`
	NEAR_HYDRANT_DISTANCE   string `json:"NEAR_HYDRANT_DISTANCE,omitempty"`
	NEAR_GT_DIST            string `json:"NEAR_GT_DIST,omitempty"`
	SERV_STAT               string `json:"SERV_STAT,omitempty"`
	FINAL_MATCH             string `json:"FINAL_MATCH,omitempty"`
	E311_SOURCE             string `json:"E311_SOURCE,omitempty"`
	E311_STATUS             string `json:"E311_STATUS,omitempty"`
	NEAR_DIST               string `json:"NEAR_DIST,omitempty"`
	W_CUSTOMERNUMBER        string `json:"W_CUSTOMERNUMBER,omitempty"`
	REF_PARCEL_UNIQUE_ID    string `json:"REF_PARCEL_UNIQUE_ID,omitempty"`
	REF_WM_UNIQUE_ID        string `json:"REF_WM_UNIQUE_ID,omitempty"`
	W_SWB_NAME              string `json:"W_SWB_NAME,omitempty"`
	NEAR_MH_FID             string `json:"NEAR_MH_FID,omitempty"`
	POLICE_BEAT             string `json:"POLICE_BEAT,omitempty"`
	PARCEL_TEST             string `json:"PARCEL_TEST,omitempty"`
	NEW_PAREL_REFERENCE_ID  string `json:"NEW_PAREL_REFERENCE_ID,omitempty"`
	NEW_PARCEL_REFERENCE_ID string `json:"NEW_PARCEL_REFERENCE_ID,omitempty"`
	E311_FINAL_LABEL        string `json:"E311_FINAL_LABEL,omitempty"`
	NEW_STATUS              string `json:"NEW_STATUS,omitempty"`
	STATUS_1                string `json:"STATUS_1,omitempty"`
	NEW_STATUS_1            string `json:"NEW_STATUS_1,omitempty"`
	E311_ADDRESS_LABEL      string `json:"E311_ADDRESS_LABEL,omitempty"`
	RANK                    string `json:"RANK,omitempty"`
	SEPARATOR               string `json:"SEPARATOR,omitempty"`
	POSTTYPE                string `json:"POSTTYPE,omitempty"`
	E311_GT_STATUS          string `json:"E311_GT_STATUS,omitempty"`
	E311_GT_ADDRESSS_TD     string `json:"E311_GT_ADDRESS_STD,omitempty"`
	NEAR_MH_ID              string `json:"NEAR_MH_ID,omitempty"`
	NEAR_MH_DIST            string `json:"NEAR_MH_DIST,omitempty"`
	NEAR_PARK_ID            string `json:"NEAR_PARK_ID,omitempty"`
	NPU_12                  string `json:"NPU_12,omitempty"`
	REF_E311_UNIQUEID       string `json:"REF_E311_UNIQUEID,omitempty"`
	W_HYDRANT_DISTANCE      string `json:"W_HYDRANT_DISTANCE,omitempty"`
	SHAPE                   string `json:"SHAPE,omitempty"`
	DAYGAR                  string `json:"DAYGAR,omitempty"`
	DAYREC                  string `json:"DAYREC,omitempty"`
	DAYYAR                  string `json:"DAYYAR,omitempty"`
	LINK                    string `json:"LINK,omitempty"`
	OWNER                   string `json:"OWNER,omitempty"`
	TAXDIST                 string `json:"TAXDIST,omitempty"`
	OWNERADDR1              string `json:"OWNERADDR1,omitempty"`
	OWNERADDR2              string `json:"OWNERADDR2,omitempty"`
	TOTASSESS               string `json:"TOTASSESS,omitempty"`
	LANDASSESS              string `json:"LANDASSESS,omitempty"`
	IMPRASSESS              string `json:"IMPRASSESS,omitempty"`
	TOTAPPR                 string `json:"TOTAPPR,omitempty"`
	LANDAPPR                string `json:"LANDAPPR,omitempty"`
	IMPRAPPR                string `json:"IMPRAPPR,omitempty"`
	LUCODE                  string `json:"LUCODE,omitempty"`
	CLASSCODE               string `json:"CLASSCODE,omitempty"`
	LIVUNITS                string `json:"LIVUNITS,omitempty"`
	LANDACRES               string `json:"LANDACRES,omitempty"`
	NBRHOOD                 string `json:"NBRHOOD,omitempty"`
	SUBDIV                  string `json:"SUBDIV,omitempty"`
	SUBDIVNUM               string `json:"SUBDIVNUM,omitempty"`
	SUBDIVLOT               string `json:"SUBDIVLOT,omitempty"`
	SUBDIVBLCK              string `json:"SUBDIVBLCK,omitempty"`
	OWNERSTATE              string `json:"OWNERSTATE,omitempty"`
	OWNERZIP                string `json:"OWNERZIP,omitempty"`
	DAYBUL                  string `json:"DAYBUL,omitempty"`
	DAYCUT                  string `json:"DAYCUT,omitempty"`
	DAYSWE                  string `json:"DAYSWE,omitempty"`
	OID                     string `json:"OID,omitempty"`
}

type Result struct {
	Records []Record `json:"records"`
}
