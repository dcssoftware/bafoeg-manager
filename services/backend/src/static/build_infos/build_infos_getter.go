package static

import "time"

func GetBuildVersion() string {
	if buildVersion == "%%APPLICATION_BUILD_VERSION%%" {
		return "0.0.0-dev"
	}
	return buildVersion
}

func GetBuildDate() string {
	if buildDate == "%%APPLICATION_BUILD_DATE%%" {
		return time.Now().String()
	}
	return buildDate
}

func GetBuildCommit() string {
	if buildCommit == "%%APPLICATION_BUILD_COMMIT%%" {
		return "0000000000000000000000000000000000000000"
	}
	return GetBuildCommit()
}
