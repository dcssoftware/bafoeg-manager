interface EaktenApplicationMappingModel {
  id: string;
  application_id: string;
  eakte_akte_id: string;
  application: EaktenApplicationMappingApplicationShortModel;
  applicant: EaktenApplicationMappingApplicantShortModel;
  school_degree: EaktenApplicationMappingDegreeShortModel;
  school: EaktenApplicationMappingSchoolShortModel;
  eakte_akte: EaktenApplicationMappingAkteShortModel;
}

interface EaktenApplicationMappingApplicationShortModel {
  id: string;
  status: EaktenApplicationMappingApplicationStatusShortModel;
}

interface EaktenApplicationMappingApplicationStatusShortModel {
  id: string;
  identifier: string;
  name: string;
}

interface EaktenApplicationMappingApplicantShortModel {
  id: string;
  firstname: string;
  lastname: string;
}

interface EaktenApplicationMappingDegreeShortModel {
  id: string;
  name: string;
}

interface EaktenApplicationMappingSchoolShortModel {
  id: string;
  name: string;
}

interface EaktenApplicationMappingAkteShortModel {
  id: string;
  aktenzeichen: string;
}