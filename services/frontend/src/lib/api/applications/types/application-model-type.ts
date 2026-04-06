export interface ApplicationModelResponseType {
    application: ApplicationModelResponseType[];
    count: number;
    maxCount: number;
}

export interface ApplicationModelResponseType {
    id: string;
    classLevel: string;
    labels: ApplicationLabelModelType[];
    applicant: ApplicationApplicantModelType;
    school: ApplicationSchoolModelType;
    assignedUser: ApplicationAssignedUserModelType;
    processingTime: ApplicationProcessingTimeModelType;
    status: ApplicationStatusModelType;
    created: Date;
    updated: Date;
}

export interface ApplicationLabelModelType {
    id: string;
    name: string;
    color: ApplicationLabelColorModelType;
}
export interface ApplicationLabelColorModelType {
    id: string;
    name: string;

    colorDark: string;
    bgColorDark: string;
    borderColorDark: string;

    colorLight: string;
    bgColorLight: string;
    borderColorLight: string;
}

export interface ApplicationApplicantModelType {
    id: string;
    firstname: string;
    lastname: string;

    address: ApplicationApplicantAddressModelType;
    trainingsAddress?: ApplicationApplicantTrainingsAddressModelType;
    contactData: ApplicationApplicantContactDataModelType;
}

export interface ApplicationApplicantAddressModelType {
    street: string;
    houseNumber: string;
    zipCode: string;
    city: string;
    country: string;
}

export interface ApplicationApplicantTrainingsAddressModelType {
    street: string;
    houseNumber: string;
    zipCode: string;
    city: string;
    country: string;
}

export interface ApplicationApplicantContactDataModelType {
    email: string;
    phone: string;
}

export interface ApplicationAssignedUserModelType {
    id: string;
    username: string;
    displayName: string;
}
export interface ApplicationSchoolModelType {
    id: string;
    name: string

    address: ApplicationSchoolAddressModelType;
    degree: ApplicationSchoolDegreeModelType;
    type: ApplicationSchoolTypeModelType;
}

export interface ApplicationSchoolDegreeModelType {
    id: string;
    name: string;
    fosBerufsabschlussRequired: boolean;
    bosBerufsqualifizierenderAbschluss: boolean;
    fachschuleBerufsschule: boolean;
}

export interface ApplicationSchoolAddressModelType {
    street: string;
    houseNumber: string;
    zipCode: string;
    city: string;
    country: string;
}

export interface ApplicationSchoolTypeModelType {
    id: string;
    typeName: string;
    typeIdentifier: string;
}

export interface ApplicationProcessingTimeModelType {
    maxValidity: number;
    remainingTimeInDays: number;
    remainingTimeInPercent: number;
    isStillLegal: boolean;
}

export interface ApplicationStatusModelType {
    id: string;
    name: string;
    identifier: string;
}




