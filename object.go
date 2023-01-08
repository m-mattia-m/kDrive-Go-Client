package kDrive

type ObjectType string

func (ot ObjectType) String() string {
	return string(ot)
}

type Cursor string

func (c Cursor) String() string {
	return string(c)
}

type FileId string

func (fId FileId) String() string {
	return string(fId)
}

type List struct {
	Result       string              `json:"result"`
	Data         []FileDirectoryList `json:"data"`
	Total        int                 `json:"total"`
	Page         int                 `json:"page"`
	Pages        int                 `json:"pages"`
	ItemsPerPage int                 `json:"items_per_page"`
	ResponseAt   int                 `json:"response_at"`
}

type FileDirectoryList struct {
	Id                     int                    `json:"id"`
	Name                   string                 `json:"name"`
	SortedName             string                 `json:"sorted_name"`
	Path                   string                 `json:"path"`
	Type                   string                 `json:"type"`
	Status                 string                 `json:"status"`
	Visibility             string                 `json:"visibility"`
	DriveId                int                    `json:"drive_id"`
	Depth                  int                    `json:"depth"`
	CreatedBy              int                    `json:"created_by"`
	CreatedAt              int                    `json:"created_at"`
	AddedAt                int                    `json:"added_at"`
	LastModifiedAt         int                    `json:"last_modified_at"`
	ParentId               int                    `json:"parent_id"`
	DeletedAt              int                    `json:"deleted_at"`
	DeletedBy              int                    `json:"deleted_by"`
	SharedRootId           int                    `json:"shared_root_id"`
	Users                  []int                  `json:"users"`
	Teams                  []int                  `json:"teams"`
	IsFavorite             bool                   `json:"is_favorite"`
	Activity               Activity               `json:"activity"`
	ShareLink              ShareLink              `json:"sharelink"`
	Capabilities           Capabilities           `json:"capabilities"`
	Lock                   Lock                   `json:"lock"`
	Categories             []FileCategory         `json:"categories"`
	Etag                   string                 `json:"etag"`
	Color                  string                 `json:"color"`
	Dropbox                Dropbox                `json:"dropbox"`
	ExternalImport         ExternalImport         `json:"external_import"`
	Parents                []Directory            `json:"parents"`
	Size                   int                    `json:"size"`
	HasThumbnail           bool                   `json:"has_thumbnail"`
	HasOnlyOffice          bool                   `json:"has_onlyoffice"`
	MimeType               string                 `json:"mime_type"`
	ExtensionType          string                 `json:"extension_type"`
	Version                Version                `json:"version"`
	ConversionCapabilities ConversionCapabilities `json:"conversion_capabilities"`
}

type DirectoryList struct {
	Id             int            `json:"id"`
	Name           string         `json:"name"`
	SortedName     string         `json:"sorted_name"`
	Path           string         `json:"path"`
	Type           string         `json:"type"`
	Status         string         `json:"status"`
	Visibility     string         `json:"visibility"`
	DriveId        int            `json:"drive_id"`
	Depth          int            `json:"depth"`
	CreatedBy      int            `json:"created_by"`
	CreatedAt      int            `json:"created_at"`
	AddedAt        int            `json:"added_at"`
	LastModifiedAt int            `json:"last_modified_at"`
	ParentId       int            `json:"parent_id"`
	DeletedAt      int            `json:"deleted_at"`
	DeletedBy      int            `json:"deleted_by"`
	SharedRootId   int            `json:"shared_root_id"`
	Users          []int          `json:"users"`
	Teams          []int          `json:"teams"`
	IsFavorite     bool           `json:"is_favorite"`
	Activity       Activity       `json:"activity"`
	ShareLink      ShareLink      `json:"sharelink"`
	Capabilities   Capabilities   `json:"capabilities"`
	Lock           Lock           `json:"lock"`
	Categories     []FileCategory `json:"categories"`
	Etag           string         `json:"etag"`
	Color          string         `json:"color"`
	Dropbox        Dropbox        `json:"dropbox"`
	ExternalImport ExternalImport `json:"external_import"`
}

type FileList struct {
	Id                     int                    `json:"id"`
	Name                   string                 `json:"name"`
	SortedName             string                 `json:"sorted_name"`
	Path                   string                 `json:"path"`
	Type                   string                 `json:"type"`
	Status                 string                 `json:"status"`
	Visibility             string                 `json:"visibility"`
	DriveId                int                    `json:"drive_id"`
	Depth                  int                    `json:"depth"`
	CreatedBy              int                    `json:"created_by"`
	CreatedAt              int                    `json:"created_at"`
	AddedAt                int                    `json:"added_at"`
	LastModifiedAt         int                    `json:"last_modified_at"`
	ParentId               int                    `json:"parent_id"`
	DeletedAt              int                    `json:"deleted_at"`
	DeletedBy              int                    `json:"deleted_by"`
	SharedRootId           int                    `json:"shared_root_id"`
	Parents                []Directory            `json:"parents"`
	Users                  []int                  `json:"users"`
	Teams                  []int                  `json:"teams"`
	IsFavorite             bool                   `json:"is_favorite"`
	Activity               Activity               `json:"activity"`
	ShareLink              ShareLink              `json:"sharelink"`
	Capabilities           Capabilities           `json:"capabilities"`
	Lock                   Lock                   `json:"lock"`
	Categories             []FileCategory         `json:"categories"`
	Etag                   string                 `json:"etag"`
	Size                   int                    `json:"size"`
	HasThumbnail           bool                   `json:"has_thumbnail"`
	HasOnlyOffice          bool                   `json:"has_onlyoffice"`
	MimeType               string                 `json:"mime_type"`
	ExtensionType          string                 `json:"extension_type"`
	Version                Version                `json:"version"`
	ConversionCapabilities ConversionCapabilities `json:"conversion_capabilities"`
}

type Directory struct {
	Result         string         `json:"result"`
	Id             int            `json:"id"`
	Name           string         `json:"name"`
	SortedName     string         `json:"sorted_name"`
	Path           string         `json:"path"`
	Type           string         `json:"type"`
	Status         string         `json:"status"`
	Visibility     string         `json:"visibility"`
	DriveId        int            `json:"drive_id"`
	Depth          int            `json:"depth"`
	CreatedBy      int            `json:"created_by"`
	CreatedAt      int            `json:"created_at"`
	AddedAt        int            `json:"added_at"`
	LastModifiedAt int            `json:"last_modified_at"`
	ParentId       int            `json:"parent_id"`
	DeletedAt      int            `json:"deleted_at"`
	DeletedBy      int            `json:"deleted_by"`
	SharedRootId   int            `json:"shared_root_id"`
	Users          []int          `json:"users"`
	Teams          []int          `json:"teams"`
	IsFavorite     bool           `json:"is_favorite"`
	Activity       Activity       `json:"activity"`
	ShareLink      ShareLink      `json:"sharelink"`
	Capabilities   Capabilities   `json:"capabilities"`
	Lock           Lock           `json:"lock"`
	Categories     []FileCategory `json:"categories"`
	Etag           string         `json:"etag"`
	Color          string         `json:"color"`
	Dropbox        Dropbox        `json:"dropbox"`
	ExternalImport ExternalImport `json:"external_import"`
}

type File struct {
	Result                 string                 `json:"result"`
	Id                     int                    `json:"id"`
	Name                   string                 `json:"name"`
	SortedName             string                 `json:"sorted_name"`
	Path                   string                 `json:"path"`
	Type                   string                 `json:"type"`
	Status                 string                 `json:"status"`
	Visibility             string                 `json:"visibility"`
	DriveId                int                    `json:"drive_id"`
	Depth                  int                    `json:"depth"`
	CreatedBy              int                    `json:"created_by"`
	CreatedAt              int                    `json:"created_at"`
	AddedAt                int                    `json:"added_at"`
	LastModifiedAt         int                    `json:"last_modified_at"`
	ParentId               int                    `json:"parent_id"`
	DeletedAt              int                    `json:"deleted_at"`
	DeletedBy              int                    `json:"deleted_by"`
	SharedRootId           int                    `json:"shared_root_id"`
	Parents                []Directory            `json:"parents"`
	Users                  []int                  `json:"users"`
	Teams                  []int                  `json:"teams"`
	IsFavorite             bool                   `json:"is_favorite"`
	Activity               Activity               `json:"activity"`
	ShareLink              ShareLink              `json:"sharelink"`
	Capabilities           Capabilities           `json:"capabilities"`
	Lock                   Lock                   `json:"lock"`
	Categories             []FileCategory         `json:"categories"`
	Etag                   string                 `json:"etag"`
	Size                   int                    `json:"size"`
	HasThumbnail           bool                   `json:"has_thumbnail"`
	HasOnlyOffice          bool                   `json:"has_onlyoffice"`
	MimeType               string                 `json:"mime_type"`
	ExtensionType          string                 `json:"extension_type"`
	Version                Version                `json:"version"`
	ConversionCapabilities ConversionCapabilities `json:"conversion_capabilities"`
}

type ExternalImport struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Status    string `json:"status"`
	Message   string `json:"message"`
	CreatedAt int    `json:"created_at"`
}

type Version struct {
	IsMultiple  bool `json:"is_multiple"`
	Number      int  `json:"number"`
	TotalSize   int  `json:"total_size"`
	KeepForever bool `json:"keep_forever"`
}

type ConversionCapabilities struct {
	WhenDownloading       bool   `json:"when_downloading"`
	DownloadExtensions    int    `json:"download_extensions"`
	WhenOnlyOfficeOpening bool   `json:"when_onlyoffice_opening"`
	OnlyOfficeExtension   string `json:"onlyoffice_extension"`
}

type Activity struct {
	LastAccessedAt int `json:"last_accessed_at"`
}

type ShareLink struct {
	Url           string `json:"url"`
	FileId        int    `json:"file_id"`
	Right         string `json:"right"`
	ValidUntil    int    `json:"valid_until"`
	CreatedBy     int    `json:"created_by"`
	CreatedAt     int    `json:"created_at"`
	UpdatedAt     int    `json:"updated_at"`
	Capabilities  int    `json:"capabilities"`
	AccessBlocked int    `json:"access_blocked"`
}

type Capabilities struct {
	CanUseFavorite     bool `json:"can_use_favorite"`
	CanBecomeSharelink bool `json:"can_become_sharelink"`
	CanUseTeam         bool `json:"can_use_team"`
	CanShow            bool `json:"can_show"`
	CanRead            bool `json:"can_read"`
	CanWrite           bool `json:"can_write"`
	CanShare           bool `json:"can_share"`
	CanLeave           bool `json:"can_leave"`
	CanDelete          bool `json:"can_delete"`
	CanRename          bool `json:"can_rename"`
	CanMove            bool `json:"can_move"`
	CanCreateDirectory bool `json:"can_create_directory"`
	CanCreateFile      bool `json:"can_create_file"`
	CanUpload          bool `json:"can_upload"`
	CanMoveInto        bool `json:"can_move_into"`
	CanBecomeDropbox   bool `json:"can_become_dropbox"`
}

type Lock struct {
	LockedAt    int    `json:"locked_at"`
	UnlockedAt  int    `json:"unlocked_at"`
	Description string `json:"description"`
	Token       string `json:"token"`
}

type Dropbox struct {
	Id             int          `json:"id"`
	Uuid           string       `json:"uuid"`
	Name           string       `json:"name"`
	Url            string       `json:"url"`
	UsersCount     int          `json:"users_count"`
	CreatedBy      int          `json:"created_by"`
	CreatedAt      int          `json:"created_at"`
	UpdatedAt      int          `json:"updated_at"`
	LastUploadedAt int          `json:"last_uploaded_at"`
	Capabilities   Capabilities `json:"capabilities"`
}

type FileCategory struct {
	CategoryId      int      `json:"category_id"`
	AddedAt         int      `json:"added_at"`
	UserValidation  string   `json:"user_validation"`
	IsGeneratedByAi bool     `json:"is_generated_by_ai"`
	UserId          int      `json:"user_id"`
	Category        Category `json:"category"`
}

type Category struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Color        string `json:"color"`
	IsPredefined bool   `json:"is_predefined"`
	CreatedBy    int    `json:"created_by"`
	CreatedAt    int    `json:"created_at"`
}

type FileStream struct {
	Name string `json:"name"`
	Type string `json:"type"`
	File []byte `json:"file"`
}
