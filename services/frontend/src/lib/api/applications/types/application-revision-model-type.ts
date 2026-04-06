export interface ApplicationRevisionModelResponseType {
    count: number;
    maxCount: number;
    revisions: ApplicationRevisionModelType[];
}

export interface ApplicationRevisionModelType {
    id: string;
    header: string;
    description: string;
    created: Date;
}
