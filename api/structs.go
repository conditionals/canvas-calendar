package api

import "time"


type CanvasUser struct {
    Url string
    Token string
}

type CanvasAssignment struct {
	Title           string    `json:"title"`
	Description     any       `json:"description"`
	SubmissionTypes string    `json:"submission_types"`
	WorkflowState   string    `json:"workflow_state"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	AllDay          bool      `json:"all_day"`
	AllDayDate      string    `json:"all_day_date"`
	ID              string    `json:"id"`
	Type            string    `json:"type"`
	Assignment      struct {
		ID                              int       `json:"id"`
		Description                     any       `json:"description"`
		DueAt                           time.Time `json:"due_at"`
		UnlockAt                        time.Time `json:"unlock_at"`
		LockAt                          any       `json:"lock_at"`
		PointsPossible                  float64   `json:"points_possible"`
		GradingType                     string    `json:"grading_type"`
		AssignmentGroupID               int       `json:"assignment_group_id"`
		GradingStandardID               any       `json:"grading_standard_id"`
		CreatedAt                       time.Time `json:"created_at"`
		UpdatedAt                       time.Time `json:"updated_at"`
		PeerReviews                     bool      `json:"peer_reviews"`
		AutomaticPeerReviews            bool      `json:"automatic_peer_reviews"`
		Position                        int       `json:"position"`
		GradeGroupStudentsIndividually  bool      `json:"grade_group_students_individually"`
		AnonymousPeerReviews            bool      `json:"anonymous_peer_reviews"`
		GroupCategoryID                 any       `json:"group_category_id"`
		PostToSis                       bool      `json:"post_to_sis"`
		ModeratedGrading                bool      `json:"moderated_grading"`
		OmitFromFinalGrade              bool      `json:"omit_from_final_grade"`
		IntraGroupPeerReviews           bool      `json:"intra_group_peer_reviews"`
		AnonymousInstructorAnnotations  bool      `json:"anonymous_instructor_annotations"`
		AnonymousGrading                bool      `json:"anonymous_grading"`
		GradersAnonymousToGraders       bool      `json:"graders_anonymous_to_graders"`
		GraderCount                     int       `json:"grader_count"`
		GraderCommentsVisibleToGraders  bool      `json:"grader_comments_visible_to_graders"`
		FinalGraderID                   any       `json:"final_grader_id"`
		GraderNamesVisibleToFinalGrader bool      `json:"grader_names_visible_to_final_grader"`
		AllowedAttempts                 int       `json:"allowed_attempts"`
		AnnotatableAttachmentID         any       `json:"annotatable_attachment_id"`
		HideInGradebook                 bool      `json:"hide_in_gradebook"`
		SecureParams                    string    `json:"secure_params"`
		LtiContextID                    string    `json:"lti_context_id"`
		CourseID                        int       `json:"course_id"`
		Name                            string    `json:"name"`
		SubmissionTypes                 []string  `json:"submission_types"`
		HasSubmittedSubmissions         bool      `json:"has_submitted_submissions"`
		DueDateRequired                 bool      `json:"due_date_required"`
		MaxNameLength                   int       `json:"max_name_length"`
		InClosedGradingPeriod           bool      `json:"in_closed_grading_period"`
		GradedSubmissionsExist          bool      `json:"graded_submissions_exist"`
		UserSubmitted                   bool      `json:"user_submitted"`
		IsQuizAssignment                bool      `json:"is_quiz_assignment"`
		CanDuplicate                    bool      `json:"can_duplicate"`
		OriginalCourseID                any       `json:"original_course_id"`
		OriginalAssignmentID            any       `json:"original_assignment_id"`
		OriginalLtiResourceLinkID       any       `json:"original_lti_resource_link_id"`
		OriginalAssignmentName          any       `json:"original_assignment_name"`
		OriginalQuizID                  any       `json:"original_quiz_id"`
		WorkflowState                   string    `json:"workflow_state"`
		ImportantDates                  bool      `json:"important_dates"`
		ExternalToolTagAttributes       struct {
			URL               string `json:"url"`
			NewTab            bool   `json:"new_tab"`
			ExternalData      any    `json:"external_data"`
			ResourceLinkID    string `json:"resource_link_id"`
			ResourceLinkTitle any    `json:"resource_link_title"`
			ContentType       string `json:"content_type"`
			ContentID         int    `json:"content_id"`
			CustomParams      any    `json:"custom_params"`
		} `json:"external_tool_tag_attributes"`
		Muted                    bool   `json:"muted"`
		HTMLURL                  string `json:"html_url"`
		URL                      any    `json:"url"`
		Published                bool   `json:"published"`
		OnlyVisibleToOverrides   bool   `json:"only_visible_to_overrides"`
		VisibleToEveryone        bool   `json:"visible_to_everyone"`
		LockedForUser            bool   `json:"locked_for_user"`
		SubmissionsDownloadURL   string `json:"submissions_download_url"`
		PostManually             bool   `json:"post_manually"`
		AnonymizeStudents        bool   `json:"anonymize_students"`
		RequireLockdownBrowser   bool   `json:"require_lockdown_browser"`
		RestrictQuantitativeData bool   `json:"restrict_quantitative_data"`
	} `json:"assignment"`
	HTMLURL        string    `json:"html_url"`
	ContextCode    string    `json:"context_code"`
	ContextName    string    `json:"context_name"`
	ContextColor   any       `json:"context_color"`
	EndAt          time.Time `json:"end_at"`
	StartAt        time.Time `json:"start_at"`
	URL            string    `json:"url"`
	ImportantDates bool      `json:"important_dates"`
}

type classes []struct {
	ID                          int       `json:"id"`
	Name                        string    `json:"name"`
	AccountID                   int       `json:"account_id"`
	UUID                        string    `json:"uuid"`
	StartAt                     time.Time `json:"start_at"`
	GradingStandardID           any       `json:"grading_standard_id"`
	IsPublic                    bool      `json:"is_public"`
	CreatedAt                   time.Time `json:"created_at"`
	CourseCode                  string    `json:"course_code"`
	DefaultView                 string    `json:"default_view"`
	RootAccountID               int       `json:"root_account_id"`
	EnrollmentTermID            int       `json:"enrollment_term_id"`
	License                     string    `json:"license"`
	GradePassbackSetting        any       `json:"grade_passback_setting"`
	EndAt                       any       `json:"end_at"`
	PublicSyllabus              bool      `json:"public_syllabus"`
	PublicSyllabusToAuth        bool      `json:"public_syllabus_to_auth"`
	StorageQuotaMb              int       `json:"storage_quota_mb"`
	IsPublicToAuthUsers         bool      `json:"is_public_to_auth_users"`
	HomeroomCourse              bool      `json:"homeroom_course"`
	CourseColor                 any       `json:"course_color"`
	FriendlyName                any       `json:"friendly_name"`
	ApplyAssignmentGroupWeights bool      `json:"apply_assignment_group_weights"`
	Calendar                    struct {
		Ics string `json:"ics"`
	} `json:"calendar"`
	TimeZone    string `json:"time_zone"`
	Blueprint   bool   `json:"blueprint"`
	Template    bool   `json:"template"`
	Enrollments []struct {
		Type                           string `json:"type"`
		Role                           string `json:"role"`
		RoleID                         int    `json:"role_id"`
		UserID                         int    `json:"user_id"`
		EnrollmentState                string `json:"enrollment_state"`
		LimitPrivilegesToCourseSection bool   `json:"limit_privileges_to_course_section"`
	} `json:"enrollments"`
	HideFinalGrades                  bool   `json:"hide_final_grades"`
	WorkflowState                    string `json:"workflow_state"`
	RestrictEnrollmentsToCourseDates bool   `json:"restrict_enrollments_to_course_dates"`
}
