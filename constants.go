package zoom

const (
	// api
	apiScheme  = "https"
	apiURI     = "api.zoom.us"
	apiVersion = "/v2"

	// paths
	createUserPath     = "/users"
	deleteUserPath     = "/users/%s"
	createMeetingPath  = "/users/%s/meetings"
	deleteMeetingPath  = "/meetings/%d"
	addUserToGroupPath = "/groups/%s/members"
	getUserPath        = "/users/%s"

	// defaults
	headerAuthorization              = "Authorization"
	headerContentType                = "Content-UserType"
	headerContentTypeApplicationJson = "application/json"
	headerValueBearer                = "Bearer %s"

	// MeetingTypeInstant is an instant meeting
	MeetingTypeInstant MeetingType = 1
	// MeetingTypeScheduled is a scheduled meeting
	MeetingTypeScheduled MeetingType = 2 // Default
	// MeetingTypeRecurringNoFixedTime is a recurring meeting with no fixed time
	MeetingTypeRecurringNoFixedTime MeetingType = 3
	// MeetingTypeRecurring is a recurring meeting with fixed time
	MeetingTypeRecurring MeetingType = 8

	// MeetingStatusWaiting is a meeting that is waiting
	MeetingStatusWaiting MeetingStatus = "waiting"
	// MeetingStatusStarted is a meeting that is started
	MeetingStatusStarted MeetingStatus = "started"
	// MeetingStatusFinished is a meeting that is finished
	MeetingStatusFinished MeetingStatus = "finished"

	// ApprovalTypeAutomaticallyApprove is an automatically approved meeting
	ApprovalTypeAutomaticallyApprove ApprovalType = 0
	// ApprovalTypeManuallyApprove is a meeting that requires manual approval
	ApprovalTypeManuallyApprove ApprovalType = 1
	// ApprovalTypeNoRegistrationRequired is a meeting that requires no registration
	ApprovalTypeNoRegistrationRequired ApprovalType = 2 // DEFAULT

	// RegistrationTypeRegisterOnce Attendees register once and can attend any of the occurrences
	RegistrationTypeRegisterOnce RegistrationType = 1
	// RegistrationTypeRegisterEachTime  Attendeed need to register for each occurrence to attend
	RegistrationTypeRegisterEachTime RegistrationType = 2
	// RegistrationTypeRegisterOnceAndChooseOccurrences Attendees register once and can choose one or more occurrences to attend
	RegistrationTypeRegisterOnceAndChooseOccurrences RegistrationType = 3

	// AudioBoth is a meeting that allows telephony and VoIP
	AudioBoth Audio = "both"
	// AudioTelephony is a meeting that is telephony only
	AudioTelephony Audio = "telephony"
	// AudioVoIP is a meeting that is VoIP only
	AudioVoIP Audio = "voip"

	// AutoRecordingLocal record on local
	AutoRecordingLocal AutoRecording = "local"
	// AutoRecordingCloud record on cloud
	AutoRecordingCloud AutoRecording = "cloud"
	// AutoRecordingNone disabled
	AutoRecordingNone AutoRecording = "none"

	// GlobalDialInNumberTypeToll toll type of number
	GlobalDialInNumberTypeToll GlobalDialInNumberType = "toll"
	// GlobalDialInNumberTypeTollFree toll free type of number
	GlobalDialInNumberTypeTollFree GlobalDialInNumberType = "tollfree"

	// RecurrenceTypeDaily daily recurrence
	RecurrenceTypeDaily RecurrenceType = 1
	// RecurrenceTypeWeekly weekly recurrence
	RecurrenceTypeWeekly RecurrenceType = 2
	// RecurrenceTypeMonthly monthly recurrence
	RecurrenceTypeMonthly RecurrenceType = 3

	// MonthlyWeekLast last week of the month
	MonthlyWeekLast MonthlyWeek = -1
	// MonthlyWeekFirst first week of the month
	MonthlyWeekFirst MonthlyWeek = 1
	// MonthlyWeekSecond second week of the month
	MonthlyWeekSecond MonthlyWeek = 2
	// MonthlyWeekThird third week of the month
	MonthlyWeekThird MonthlyWeek = 3
	// MonthlyWeekFourth fourth week of the month
	MonthlyWeekFourth MonthlyWeek = 4
	//

	// WeekDaySunday Sunday
	WeekDaySunday WeekDay = 1
	// WeekDayMonday Monday
	WeekDayMonday WeekDay = 2
	// WeekDayTuesday Tuesday
	WeekDayTuesday WeekDay = 3
	// WeekDayWednesday Wednesday
	WeekDayWednesday WeekDay = 4
	// WeekDayThursday Thursday
	WeekDayThursday WeekDay = 5
	// WeekDayFriday Friday
	WeekDayFriday WeekDay = 6
	// WeekDaySaturday Saturday
	WeekDaySaturday WeekDay = 7

	// create user actions
	UserCreateActionCreate     UserCreateAction = "create"
	UserCreateActionAutoCreate UserCreateAction = "autoCreate"
	UserCreateActionCustCreate UserCreateAction = "custCreate"
	UserCreateActionSsoCreate  UserCreateAction = "ssoCreate"

	// user types
	UserTypeBasic    UserType = 1
	UserTypeLicensed UserType = 2
	UserTypeOnPrem   UserType = 3

	// user login types
	Facebook UserLoginType = 0   // Facebook user login type
	Google   UserLoginType = 1   // Google user login type
	API      UserLoginType = 99  // API user login type
	Zoom     UserLoginType = 100 // Zoom user login type
	SSO      UserLoginType = 101 // SSO single sign on user login type

	// error codes
	ErrorCodeUserNotFound = 1005
)
