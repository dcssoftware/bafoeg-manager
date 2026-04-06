import type {
    ApplicationSchoolModelType,
    ApplicationAssignedUserModelType,
    ApplicationProcessingTimeModelType
} from "./application-model-type"

export interface ApplicationShortModelsResponseType {
    application: ApplicationShortModelResponseType[];
    count: number;
    maxCount: number;
}

export interface ApplicationShortModelResponseType {
    id: string;
    applicant: ApplicationApplicantShortModelType;
    school: ApplicationSchoolModelType;
    assignedUser: ApplicationAssignedUserModelType;
    processingTime: ApplicationProcessingTimeModelType;
    created: Date;
    updated: Date;
}

export interface ApplicationApplicantShortModelType {
    id: string;
    firstname: string;
    lastname: string;
}
