/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

/*
 * The sample smart contract for documentation topic:
 * Writing Your First Blockchain m
 */

 package main

 /* Imports
  * 4 utility libraries for formatting, handling bytes, reading and writing JSON, and string manipulation
  * 2 specific Hyperledger Fabric specific libraries for Smart Contracts
  */
 import (
	 "bytes"
	 "strings"
	 "encoding/json"
	 "encoding/base64"
	 "fmt"
	 "github.com/hyperledger/fabric/core/chaincode/shim"
	 sc "github.com/hyperledger/fabric/protos/peer"
 )
 
 // Define the Smart Contract structure
 type SmartContract struct {
 }
 
 // Define the car structure, with 4 properties.  Structure tags are used by encoding/json library
 type Certificate struct {
	    ID_NUMBER string `json:"ID_NUMBER"`	
            CIRCLE_NO string `json:"CIRCLE_NO"`
            WARD_NO string `json:"WARD_NO"`
            TYPE_OF_HOSPITAL string `json:"TYPE_OF_HOSPITAL"`
            HOSPITAL_NAME string `json:"HOSPITAL_NAME"`
            HOSPITAL_ADDRESS string `json:"HOSPITAL_ADDRESS"`
            ATTENTION_TYPE string `json:"ATTENTION_TYPE"`
            INFORMANT_NAME string `json:"INFORMANT_NAME"`
			DATE_OF_BIRTH string `json:"DATE_OF_BIRTH"`
            CHILD_FIRST_NAME string `json:"CHILD_FIRST_NAME"`
            GENDER string `json:"GENDER"`
            BABY_WEIGHT string `json:"BABY_WEIGHT"`
            BIRTH_ORDER string `json:"BIRTH_ORDER"`
            NO_OF_BABIES string `json:"NO_OF_BABIES"`
            BIRTH_TYPE string `json:"BIRTH_TYPE"`
            STILL_DEATH_CAUSE string `json:"STILL_DEATH_CAUSE"`
            STILL_DEATH_CAUSE_DETAILS string `json:"STILL_DEATH_CAUSE_DETAILS"`
            MOTHER_FIRST_NAME string `json:"MOTHER_FIRST_NAME"`
            MOTHER_LITERACY string `json:"MOTHER_LITERACY"`
            MOTHER_OCCUPATION string `json:"MOTHER_OCCUPATION"`
            MOTHER_RELIGION string `json:"MOTHER_RELIGION"`
            MOTHER_NATIONALITY string `json:"MOTHER_NATIONALITY"`
            MOTHER_AGE_AT_MARRIAGE string `json:"MOTHER_AGE_AT_MARRIAGE"`
            MOTHER_AGE_AT_DELIVERY string `json:"MOTHER_AGE_AT_DELIVERY"`
            PREGNANT_DURATION string `json:"PREGNANT_DURATION"`
            DELIVERY_METHOD string `json:"DELIVERY_METHOD"`
            MOTHER_MOBILE string `json:"MOTHER_MOBILE"`
            FATHER_NAME string `json:"FATHER_NAME"`
            FATHER_LITERACY string `json:"FATHER_LITERACY"`
            FATHER_OCCUPATION string `json:"FATHER_OCCUPATION"`
            FATHER_RELIGION string `json:"FATHER_RELIGION"` 
            FATHER_NATIONALITY string `json:"FATHER_NATIONALITY"`
            FATHER_MOBILE string `json:"FATHER_MOBILE"`
            RESIDENCE_ADDRESS string `json:"RESIDENCE_ADDRESS"`
            REG_NO string `json:"REG_NO"`
            REG_DATE string `json:"REG_DATE"`
            ACK_NO string `json:"ACK_NO"`
            ACK_DATE string `json:"ACK_DATE"`
            REMARKS string `json:"REMARKS"`
            PRINT_STATUS string `json:"PRINT_STATUS"`
            FIRST_SIGNED_BY string `json:"FIRST_SIGNED_BY"`
            FIRST_SIGNED_DATE string `json:"FIRST_SIGNED_DATE"`
            APPROVAL_STATUS string `json:"APPROVAL_STATUS"`
            REF_DATE string `json:"REF_DATE"`
            CERT_PRINT string `json:"CERT_PRINT"`
            CERT_PRINT_DATE string `json:"CERT_PRINT_DATE"`
            EIDNUMBER string `json:"EIDNUMBER"`
            ANYDEFECT string `json:"ANYDEFECT"`
            DEFECTDESC string `json:"DEFECTDESC"`
            MOTHER_UID_NUMBER string `json:"MOTHER_UID_NUMBER"`
            FATHER_UID_NUMBER string `json:"FATHER_UID_NUMBER"`
            MOTHER_KCR_KITNO string `json:"MOTHER_KCR_KITNO"`
            ZERO_DOSE_IMMNZN string `json:"ZERO_DOSE_IMMNZN"`
            CERTIFICATE_NUMBER string `json:"CERTIFICATE_NUMBER"`
            NO_OF_COPIES string `json:"NO_OF_COPIES"`
            FEE_COLLECTED string `json:"FEE_COLLECTED"`
            RES_DISTRICT string `json:"RES_DISTRICT"` 
            RES_PINCODE string `json:"RES_PINCODE"`
            MOBILE_NO string `json:"MOBILE_NO"`
            PHONE_NO string `json:"PHONE_NO"`
            EMAIL_ID string `json:"EMAIL_ID"`
            CHILD_NAME_CORR string `json:"CHILD_NAME_CORR"`
            FATHER_NAME_CORR string `json:"FATHER_NAME_CORR"`
            MOTHER_NAME_CORR string `json:"MOTHER_NAME_CORR"`
            APPLIED_BY string `json:"APPLIED_BY"`
            RELATION string `json:"RELATION"`
            DOB_CORR string `json:"DOB_CORR"`
            GENDER_CORR string `json:"GENDER_CORR"`
            BIRTH_PLACE_CORR string `json:"BIRTH_PLACE_CORR"`
            PERM_ADDRESS string `json:"PERM_ADDRESS"`
            C_NAME_CORRECTED_YES_NO bool `json:"C_NAME_CORRECTED_YES_NO"`
            F_NAME_CORRECTED_YES_NO bool `json:"F_NAME_CORRECTED_YES_NO"`
            M_NAME_CORRECTED_YES_NO bool `json:"M_NAME_CORRECTED_YES_NO"`
            DOB_CORRECTED_YES_NO bool `json:"DOB_CORRECTED_YES_NO"`
            GENDER_CORRECTED_YES_NO bool `json:"GENDER_CORRECTED_YES_NO"`
            B_PLACE_CORRECTED_YES_NO bool `json:"B_PLACE_CORRECTED_YES_NO"`
            DIGITALLY_SIGNED bool `json:"DIGITALLY_SIGNED"`
            REQ_APPROVAL_STATUS string `json:"REQ_APPROVAL_STATUS"`
            REQ_SIGNED_BY string `json:"REQ_SIGNED_BY"`
            REQ_SIGNED_DATE string `json:"REQ_SIGNED_DATE"`
            CSC_USERID string `json:"CSC_USERID"`
            CIRCLE_ACCEPTED bool `json:"CIRCLE_ACCEPTED"`
            CIRCLE_ACCEPTED_DATE string `json:"CIRCLE_ACCEPTED_DATE"`
            OLD_CERTIFICATE_NUMBER string `json:"OLD_CERTIFICATE_NUMBER"` 
            IS_MODIFICATION_REQUEST bool `json:"IS_MODIFICATION_REQUEST"`
            IS_REQUEST_DIGITALLY_SIGNED bool `json:"IS_REQUEST_DIGITALLY_SIGNED"`
 }
 
 /*
  * The Init method is called when the Smart Contract "fabcar" is instantiated by the blockchain network
  * Best practice is to have any Ledger initialization in separate function -- see initLedger()
  */
 func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	 return shim.Success(nil)
 }
 
 /*
  * The Invoke method is called as a result of an m request to run the Smart Contract "fabcar"
  * The calling m program has also specified the particular smart contract function to be called, with arguments
  */
 func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {
 
	 // Retrieve the requested Smart Contract function and arguments
	 function, args := APIstub.GetFunctionAndParameters()
	 // Route to the appropriate handler function to interact with the ledger appropriately
	 if function == "queryCertificate" {
		 return s.queryCertificate(APIstub, args)
	 } else if function == "initLedger" {
		 return s.initLedger(APIstub)
	 } else if function == "register" {
		 return s.register(APIstub)
	 } else if function == "queryAllCertificates" {
		 return s.queryAllCertificates(APIstub)
	 } else if function == "registrationApproval" {
		 return s.registrationApproval(APIstub)
	 } else if function == "digitalSignature" {
		return s.digitalSignature(APIstub)
	 } else if function == "changeRequest" {
		return s.changeRequest(APIstub)
	 } else if function == "approveOrRejectRequest" {
		return s.approveOrRejectRequest(APIstub)
	 } else if function == "digitalSignatureForChangeRequest" {
		return s.digitalSignatureForChangeRequest(APIstub)
	 }   
 
	 return shim.Error("Invalid Smart Contract function name.")
 }
 
 func (s *SmartContract) queryCertificate(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
 
	 if len(args) != 1 {
		 return shim.Error("Incorrect number of arguments. Expecting 1")
	 }
 
	 carAsBytes, _ := APIstub.GetState(args[0])
	 return shim.Success(carAsBytes)
 }
 
 func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	 certificates := []Certificate{
		Certificate{
		CIRCLE_NO: "1267",
		WARD_NO: "1c",
		TYPE_OF_HOSPITAL: "PVT Hospital",
		HOSPITAL_NAME: "Neelima Hospitals",
		HOSPITAL_ADDRESS: "Jubilee Hills",
		ATTENTION_TYPE: "PVT",
		INFORMANT_NAME: "DR. SUDHEER",
		DATE_OF_BIRTH: "2019/05/22",
		CHILD_FIRST_NAME: "SUJITH K",
		GENDER: "MALE",
		BABY_WEIGHT: "2",
		BIRTH_ORDER: "5",
		NO_OF_BABIES: "1",
		BIRTH_TYPE: "NORMAL",
		STILL_DEATH_CAUSE: "",
		STILL_DEATH_CAUSE_DETAILS: "",
		MOTHER_FIRST_NAME: "SUNITHA K",
		MOTHER_LITERACY: "BSC",
		MOTHER_OCCUPATION: "HOUSE WIFE",
		MOTHER_RELIGION: "HINDU",
		MOTHER_NATIONALITY: "INDIAN",
		MOTHER_AGE_AT_MARRIAGE: "23",
		MOTHER_AGE_AT_DELIVERY: "25",
		PREGNANT_DURATION: "10",
		DELIVERY_METHOD: "NORMAL",
		MOTHER_MOBILE: "9775566223",
		FATHER_NAME: "KUMAR K",
		FATHER_LITERACY: "MSC",
		FATHER_OCCUPATION: "MANAGER",
		FATHER_RELIGION: "HINDU",
		FATHER_NATIONALITY: "INDIAN",
		FATHER_MOBILE: "9412586321",
		RESIDENCE_ADDRESS: "MADHAPUR",
		REG_NO: "",
		REG_DATE: "",
		ACK_NO: "",
		ACK_DATE: "",
		REMARKS: "",
		PRINT_STATUS: "",
		FIRST_SIGNED_BY: "",
		FIRST_SIGNED_DATE: "",
		APPROVAL_STATUS: "",
		REF_DATE: "",
		CERT_PRINT: "",
		CERT_PRINT_DATE: "",
		EIDNUMBER: "",
		ANYDEFECT: "",
		DEFECTDESC: "",
		MOTHER_UID_NUMBER: "",
		FATHER_UID_NUMBER: "",
		MOTHER_KCR_KITNO: "",
		ZERO_DOSE_IMMNZN: "",
		CERTIFICATE_NUMBER: "",
		NO_OF_COPIES: "",
		FEE_COLLECTED: "",
		RES_DISTRICT: "",
		RES_PINCODE: "",
		MOBILE_NO: "",
		PHONE_NO: "",
		EMAIL_ID: "",
		CHILD_NAME_CORR: "",
		FATHER_NAME_CORR: "",
		MOTHER_NAME_CORR: "",
		APPLIED_BY: "",
		RELATION: "",
		DOB_CORR: "",
		GENDER_CORR: "",
		BIRTH_PLACE_CORR: "",
		PERM_ADDRESS: "",
		C_NAME_CORRECTED_YES_NO: false,
		F_NAME_CORRECTED_YES_NO: false,
		M_NAME_CORRECTED_YES_NO: false,
		DOB_CORRECTED_YES_NO: false,
		GENDER_CORRECTED_YES_NO: false,
		B_PLACE_CORRECTED_YES_NO: false,
		DIGITALLY_SIGNED: false,
		REQ_APPROVAL_STATUS: "",
		REQ_SIGNED_BY: "",
		REQ_SIGNED_DATE: "",
		CSC_USERID: "",
		CIRCLE_ACCEPTED: false,
		CIRCLE_ACCEPTED_DATE: "",
		OLD_CERTIFICATE_NUMBER: "",
		IS_MODIFICATION_REQUEST: false,
		IS_REQUEST_DIGITALLY_SIGNED: false},
	 }
 
	 i := 0
	 for i < len(certificates) {
		 fmt.Println("i is ", i)
		 certificateAsBytes, _ := json.Marshal(certificates[i])
		 APIstub.PutState("ID-0", certificateAsBytes)
		 fmt.Println("Added", certificates[i])
		 i = i + 1
	 }
	 return shim.Success(nil)
 }
 
 func (s *SmartContract) register(APIstub shim.ChaincodeStubInterface) sc.Response {
	function, args := APIstub.GetFunctionAndParameters()
	fmt.Println("Funciton:",function)
	fmt.Println("Args:", args);
	fmt.Println("ArgsLength:", len(args));
	justString := strings.Join(args,"")
	//  if len(args) != 1 {
	// 	 return shim.Error("Incorrect number of arguments. Expecting 1")
	//  }
	 
	 var m Certificate
	 decoded, errDecode := base64.StdEncoding.DecodeString(justString)
	 if errDecode !=nil {
		 return shim.Error("Incorrect encoded String")
	 }
	 err := json.Unmarshal(decoded, &m)

	 fmt.Println("Certificate", &m)
	 if err != nil {
		return shim.Error(err.Error())
	}
	 var applicationToRegister = Certificate{
	 CIRCLE_NO: m.CIRCLE_NO,
	 WARD_NO: m.WARD_NO,
	 TYPE_OF_HOSPITAL: m.TYPE_OF_HOSPITAL,
	 HOSPITAL_NAME: m.HOSPITAL_NAME,
	 HOSPITAL_ADDRESS: m.HOSPITAL_ADDRESS,
	 ATTENTION_TYPE: m.ATTENTION_TYPE,
	 INFORMANT_NAME: m.INFORMANT_NAME,
	 DATE_OF_BIRTH: m.DATE_OF_BIRTH,
	 CHILD_FIRST_NAME: m.CHILD_FIRST_NAME,
	 GENDER: m.GENDER,
	 BABY_WEIGHT: m.BABY_WEIGHT,
	 BIRTH_ORDER: m.BIRTH_ORDER,
	 NO_OF_BABIES: m.NO_OF_BABIES,
	 BIRTH_TYPE: m.BIRTH_TYPE,
	 STILL_DEATH_CAUSE: m.STILL_DEATH_CAUSE,
	 STILL_DEATH_CAUSE_DETAILS: m.STILL_DEATH_CAUSE_DETAILS,
	 MOTHER_FIRST_NAME: m.MOTHER_FIRST_NAME,
	 MOTHER_LITERACY: m.MOTHER_LITERACY,
	 MOTHER_OCCUPATION: m.MOTHER_OCCUPATION,
	 MOTHER_RELIGION: m.MOTHER_RELIGION,
	 MOTHER_NATIONALITY: m.MOTHER_NATIONALITY,
	 MOTHER_AGE_AT_MARRIAGE: m.MOTHER_AGE_AT_MARRIAGE,
	 MOTHER_AGE_AT_DELIVERY: m.MOTHER_AGE_AT_DELIVERY,
	 PREGNANT_DURATION: m.PREGNANT_DURATION,
	 DELIVERY_METHOD: m.DELIVERY_METHOD,
	 MOTHER_MOBILE: m.MOTHER_MOBILE,
	 FATHER_NAME: m.FATHER_NAME,
	 FATHER_LITERACY: m.FATHER_LITERACY,
	 FATHER_OCCUPATION: m.FATHER_OCCUPATION,
	 FATHER_RELIGION: m.FATHER_RELIGION,
	 FATHER_NATIONALITY: m.FATHER_NATIONALITY,
	 FATHER_MOBILE: m.FATHER_MOBILE,
	 RESIDENCE_ADDRESS: m.RESIDENCE_ADDRESS,
	 REG_NO: m.REG_NO,
	 REG_DATE: m.REG_DATE,
	 ACK_NO: m.ACK_NO,
	 ACK_DATE: m.ACK_DATE,
	 REMARKS: m.REMARKS,
	 PRINT_STATUS: m.PRINT_STATUS,
	 FIRST_SIGNED_BY: m.FIRST_SIGNED_BY,
	 FIRST_SIGNED_DATE: m.FIRST_SIGNED_DATE,
	 APPROVAL_STATUS: m.APPROVAL_STATUS,
	 REF_DATE: m.REF_DATE,
	 CERT_PRINT: m.CERT_PRINT,
	 CERT_PRINT_DATE: m.CERT_PRINT_DATE,
	 EIDNUMBER: m.EIDNUMBER,
	 ANYDEFECT: m.ANYDEFECT,
	 DEFECTDESC: m.DEFECTDESC,
	 MOTHER_UID_NUMBER: m.MOTHER_UID_NUMBER,
	 FATHER_UID_NUMBER: m.FATHER_UID_NUMBER,
	 MOTHER_KCR_KITNO: m.MOTHER_KCR_KITNO,
	 ZERO_DOSE_IMMNZN: m.ZERO_DOSE_IMMNZN,
	 CERTIFICATE_NUMBER: "",
	 NO_OF_COPIES: "",
	 FEE_COLLECTED: "",
	 RES_DISTRICT: "",
	 RES_PINCODE: "",
	 MOBILE_NO: "",
	 PHONE_NO: "",
	 EMAIL_ID: "",
	 CHILD_NAME_CORR: "",
	 FATHER_NAME_CORR: "",
	 MOTHER_NAME_CORR: "",
	 APPLIED_BY: "",
	 RELATION: "",
	 DOB_CORR: "",
	 GENDER_CORR: "",
	 BIRTH_PLACE_CORR: "",
	 PERM_ADDRESS: "",
	 C_NAME_CORRECTED_YES_NO: false,
	 F_NAME_CORRECTED_YES_NO: false,
	 M_NAME_CORRECTED_YES_NO: false,
	 DOB_CORRECTED_YES_NO: false,
	 GENDER_CORRECTED_YES_NO: false,
	 B_PLACE_CORRECTED_YES_NO: false,
	 DIGITALLY_SIGNED: false,
	 REQ_APPROVAL_STATUS: "",
	 REQ_SIGNED_BY: "",
	 REQ_SIGNED_DATE: "",
	 IS_REQUEST_DIGITALLY_SIGNED: false,
	 CSC_USERID: "",
	 CIRCLE_ACCEPTED: false,
	 CIRCLE_ACCEPTED_DATE: "",
	 OLD_CERTIFICATE_NUMBER: "",
	 IS_MODIFICATION_REQUEST: false}
 
	 certificateAsBytes, _ := json.Marshal(applicationToRegister)
	 APIstub.PutState(m.ID_NUMBER, certificateAsBytes)
 
	 return shim.Success(nil)
 }
 
 func (s *SmartContract) queryAllCertificates(APIstub shim.ChaincodeStubInterface) sc.Response {
 
	 startKey := "ID-1"
	 endKey := "ID-99999999999999"
 
	 resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	 if err != nil {
		 return shim.Error(err.Error())
	 }
	 defer resultsIterator.Close()
 
	 // buffer is a JSON array containing QueryResults
	 var buffer bytes.Buffer
	 buffer.WriteString("[")
 
	 bArrayMemberAlreadyWritten := false
	 for resultsIterator.HasNext() {
		 queryResponse, err := resultsIterator.Next()
		 if err != nil {
			 return shim.Error(err.Error())
		 }
		 // Add a comma before array members, suppress it for the first array member
		 if bArrayMemberAlreadyWritten == true {
			 buffer.WriteString(",")
		 }
		 buffer.WriteString("{\"Key\":")
		 buffer.WriteString("\"")
		 buffer.WriteString(queryResponse.Key)
		 buffer.WriteString("\"")
 
		 buffer.WriteString(", \"Record\":")
		 // Record is a JSON object, so we write as-is
		 buffer.WriteString(string(queryResponse.Value))
		 buffer.WriteString("}")
		 bArrayMemberAlreadyWritten = true
	 }
	 buffer.WriteString("]")
 
	 fmt.Printf("- queryAllCerticates:\n%s\n", buffer.String())
 
	 return shim.Success(buffer.Bytes())
 }
 
 func (s *SmartContract) registrationApproval(APIstub shim.ChaincodeStubInterface) sc.Response {
	function, args := APIstub.GetFunctionAndParameters()
	fmt.Println("Funciton:",function)
	fmt.Println("Args:", args);
	fmt.Println("ArgsLength:", len(args));
	justString := strings.Join(args,"")
	//  if len(args) != 1 {
	// 	 return shim.Error("Incorrect number of arguments. Expecting 1")
	//  }
	 
	 var m Certificate
	 decoded, errDecode := base64.StdEncoding.DecodeString(justString)
	 if errDecode !=nil {
		 return shim.Error("Incorrect encoded String")
	 }

	 err := json.Unmarshal(decoded, &m)
	 fmt.Println("m", &m)
	 if err != nil {
		return shim.Error(err.Error())
	}
 
	 certificateAsBytes, _ := APIstub.GetState(m.ID_NUMBER)
	 certificate := Certificate{}
 
	 json.Unmarshal(certificateAsBytes, &certificate)
	 certificate.REG_NO = m.REG_NO;
     certificate.REG_DATE = m.REG_DATE;
     certificate.APPROVAL_STATUS = m.APPROVAL_STATUS;
 
	 certificateAsBytes, _ = json.Marshal(certificate)
	 APIstub.PutState(m.ID_NUMBER, certificateAsBytes)
 
	 return shim.Success(nil)
 }


 func (s *SmartContract) digitalSignature(APIstub shim.ChaincodeStubInterface) sc.Response {
	function, args := APIstub.GetFunctionAndParameters()
	fmt.Println("Funciton:",function)
	fmt.Println("Args:", args);
	fmt.Println("ArgsLength:", len(args));
	justString := strings.Join(args,"")
	//  if len(args) != 1 {
	// 	 return shim.Error("Incorrect number of arguments. Expecting 1")
	//  }
	 
	 var m Certificate
	 decoded, errDecode := base64.StdEncoding.DecodeString(justString)
	 if errDecode !=nil {
		 return shim.Error("Incorrect encoded String")
	 }

	 err := json.Unmarshal(decoded, &m)
	 fmt.Println("m", &m)
	 if err != nil {
		return shim.Error(err.Error())
	}
 

	certificateAsBytes, _ := APIstub.GetState(m.ID_NUMBER)
	certificate := Certificate{}

	json.Unmarshal(certificateAsBytes, &certificate)
	certificate.FIRST_SIGNED_BY = m.FIRST_SIGNED_BY;
	certificate.FIRST_SIGNED_DATE = m.FIRST_SIGNED_DATE;
	certificate.DIGITALLY_SIGNED = true;
	certificate.CERTIFICATE_NUMBER = m.CERTIFICATE_NUMBER;

	certificateAsBytes, _ = json.Marshal(certificate)
	APIstub.PutState(m.ID_NUMBER, certificateAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) changeRequest(APIstub shim.ChaincodeStubInterface) sc.Response {
	function, args := APIstub.GetFunctionAndParameters()
	fmt.Println("Funciton:",function)
	fmt.Println("Args:", args);
	fmt.Println("ArgsLength:", len(args));
	justString := strings.Join(args,"")
	//  if len(args) != 1 {
	// 	 return shim.Error("Incorrect number of arguments. Expecting 1")
	//  }
	 
	 var m Certificate
	 decoded, errDecode := base64.StdEncoding.DecodeString(justString)
	 if errDecode !=nil {
		 return shim.Error("Incorrect encoded String")
	 }

	 err := json.Unmarshal(decoded, &m)
	 fmt.Println("m", &m)
	 if err != nil {
		return shim.Error(err.Error())
	}
 

	certificateAsBytes, _ := APIstub.GetState(m.ID_NUMBER)
	certificate := Certificate{}

	json.Unmarshal(certificateAsBytes, &certificate)
	if certificate.CERTIFICATE_NUMBER != "" {
		certificate.NO_OF_COPIES = m.NO_OF_COPIES;
		certificate.CHILD_NAME_CORR = m.CHILD_NAME_CORR;
		certificate.FATHER_NAME_CORR = m.FATHER_NAME_CORR;
		certificate.MOTHER_NAME_CORR = m.MOTHER_NAME_CORR;
		certificate.GENDER_CORR = m.GENDER_CORR;
		certificate.DOB_CORR = m.DOB_CORR;
		certificate.BIRTH_PLACE_CORR = m.BIRTH_PLACE_CORR;
		certificate.PERM_ADDRESS = m.PERM_ADDRESS;
		certificate.RESIDENCE_ADDRESS = m.RESIDENCE_ADDRESS;
		certificate.RES_DISTRICT = m.RES_DISTRICT;
		certificate.RES_PINCODE = m.RES_PINCODE;
		certificate.EMAIL_ID = m.EMAIL_ID;
		certificate.APPLIED_BY = m.APPLIED_BY;
		certificate.RELATION = m.RELATION;
		certificate.MOBILE_NO = m.MOBILE_NO;
		certificate.PHONE_NO = m.PHONE_NO;
		certificate.FEE_COLLECTED = m.FEE_COLLECTED;
		certificate.IS_MODIFICATION_REQUEST = true;
	}

	certificateAsBytes, _ = json.Marshal(certificate)
	APIstub.PutState(m.ID_NUMBER, certificateAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) approveOrRejectRequest(APIstub shim.ChaincodeStubInterface) sc.Response {
	function, args := APIstub.GetFunctionAndParameters()
	fmt.Println("Funciton:",function)
	fmt.Println("Args:", args);
	fmt.Println("ArgsLength:", len(args));
	justString := strings.Join(args,"")
	//  if len(args) != 1 {
	// 	 return shim.Error("Incorrect number of arguments. Expecting 1")
	//  }
	 
	 var m Certificate
	 decoded, errDecode := base64.StdEncoding.DecodeString(justString)
	 if errDecode !=nil {
		 return shim.Error("Incorrect encoded String")
	 }

	 err := json.Unmarshal(decoded, &m)
	 fmt.Println("m", &m)
	 if err != nil {
		return shim.Error(err.Error())
	}
 
	certificateAsBytes, _ := APIstub.GetState(m.ID_NUMBER)
	certificate := Certificate{}

	json.Unmarshal(certificateAsBytes, &certificate)
	if certificate.CHILD_NAME_CORR != "" {
		certificate.C_NAME_CORRECTED_YES_NO = true
		certificate.CHILD_FIRST_NAME= certificate.CHILD_NAME_CORR
	}
	if certificate.MOTHER_NAME_CORR != "" {
		certificate.M_NAME_CORRECTED_YES_NO = true 
		certificate.MOTHER_FIRST_NAME = certificate.MOTHER_NAME_CORR
	}
	if certificate.FATHER_NAME_CORR != "" {
		certificate.F_NAME_CORRECTED_YES_NO = true
		certificate.FATHER_NAME = certificate.FATHER_NAME_CORR
	}
	if certificate.DOB_CORR != "" {
		certificate.DOB_CORRECTED_YES_NO = true
		certificate.DATE_OF_BIRTH = certificate.DOB_CORR
	}
	if certificate.BIRTH_PLACE_CORR != "" {
		certificate.B_PLACE_CORRECTED_YES_NO = true
		certificate.HOSPITAL_NAME = certificate.BIRTH_PLACE_CORR
	}
	if certificate.GENDER_CORR != "" {
		certificate.GENDER_CORRECTED_YES_NO = true
		certificate.GENDER = certificate.GENDER_CORR
	}
	certificate.REQ_APPROVAL_STATUS = m.REQ_APPROVAL_STATUS;

	certificateAsBytes, _ = json.Marshal(certificate)
	APIstub.PutState(m.ID_NUMBER, certificateAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) digitalSignatureForChangeRequest(APIstub shim.ChaincodeStubInterface) sc.Response {
	function, args := APIstub.GetFunctionAndParameters()
	fmt.Println("Funciton:",function)
	fmt.Println("Args:", args);
	fmt.Println("ArgsLength:", len(args));
	justString := strings.Join(args,"")
	//  if len(args) != 1 {
	// 	 return shim.Error("Incorrect number of arguments. Expecting 1")
	//  }
	 
	 var m Certificate
	 decoded, errDecode := base64.StdEncoding.DecodeString(justString)
	 if errDecode !=nil {
		 return shim.Error("Incorrect encoded String")
	 }

	 err := json.Unmarshal(decoded, &m)
	 fmt.Println("m", &m)
	 if err != nil {
		return shim.Error(err.Error())
	}
 

	certificateAsBytes, _ := APIstub.GetState(m.ID_NUMBER)
	certificate := Certificate{}

	json.Unmarshal(certificateAsBytes, &certificate)
	certificate.REQ_SIGNED_BY = m.REQ_SIGNED_BY;
    certificate.REQ_SIGNED_DATE = m.REQ_SIGNED_DATE;
    certificate.IS_REQUEST_DIGITALLY_SIGNED = true;
    certificate.OLD_CERTIFICATE_NUMBER = certificate.CERTIFICATE_NUMBER;
    certificate.CERTIFICATE_NUMBER = m.CERTIFICATE_NUMBER;
	certificateAsBytes, _ = json.Marshal(certificate)
	APIstub.PutState(m.ID_NUMBER, certificateAsBytes)

	return shim.Success(nil)
}
 
 // The main function is only relevant in unit test mode. Only included here for completeness.
 func main() {
 
	 // Create a new Smart Contract
	 err := shim.Start(new(SmartContract))
	 if err != nil {
		 fmt.Printf("Error creating new Smart Contract: %s", err)
	 }
 }
 
