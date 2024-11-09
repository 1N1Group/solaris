package v2

import "net/http"

type CollectionBackupDetails struct {
	BackupID             int        `json:"backupId,omitempty"`
	IndexVersion         string     `json:"indexVersion,omitempty"`
	StartTime            string     `json:"startTime,omitempty"`
	EndTime              string     `json:"endTime,omitempty"`
	IndexFileCount       string     `json:"indexFileCount,omitempty"`
	IndexSizeMB          int        `json:"indexSizeMB,omitempty"`
	ShardBackupIds       Properties `json:"shardBackupIds,omitempty"`
	CollectionAlias      string     `json:"collectionAlias,omitempty"`
	ExtraProperties      Properties `json:"extraProperties,omitempty"`
	CollectionConfigName string     `json:"collection.configName,omitempty"`
}

type ErrorInfo struct {
	Metadata ErrorMetadata `json:"metadata,omitempty"`
	Details  []Properties  `json:"details,omitempty"`
	Msg      string        `json:"msg,omitempty"`
	Trace    string        `json:"trace,omitempty"`
	Code     int           `json:"code,omitempty"`
}

type ErrorMetadata struct {
	ErrorClass     string `json:"error-class,omitempty"`
	RootErrorClass string `json:"root-error-class,omitempty"`
}

type Properties map[string]interface{}

type PurgeUnusedStats struct {
	NumBackupIds      int `json:"numBackupIds,omitempty"`
	NumShardBackupIds int `json:"numShardBackupIds,omitempty"`
	NumIndexFiles     int `json:"numIndexFiles,omitempty"`
}

type Response struct {
	HttpResponse   *http.Response
	Aliases        Properties                     `json:"aliases,omitempty"`
	Backups        []CollectionBackupDetails      `json:"backups,omitempty"`
	Collection     string                         `json:"collection,omitempty"` // The name of the collection.
	Collections    []string                       `json:"collections,omitempty"`
	CommitName     string                         `json:"commitName,omitempty"` // The name of the created | deleted snapshot.
	ConfigSets     []string                       `json:"configSets,omitempty"`
	Core           string                         `json:"core,omitempty"`     // The name of the core.
	CoreName       string                         `json:"coreName,omitempty"` // The name of the core.
	Deleted        PurgeUnusedStats               `json:"deleted,omitempty"`
	Error          ErrorInfo                      `json:"error,omitempty"`
	Failure        interface{}                    `json:"failure,omitempty"`
	File           string                         `json:"file,omitempty"`
	Files          []string                       `json:"files,omitempty"`         // The list of index filenames contained within the created snapshot.
	FollowAliases  bool                           `json:"followAliases,omitempty"` // A flag that treats the collName parameter as a collection alias.
	Generation     int                            `json:"generation,omitempty"`    // The generation value for the created snapshot.
	IndexDirPath   string                         `json:"indexDirPath,omitempty"`  // The path to the directory containing the index files.
	Key            string                         `json:"key,omitempty"`           // The public key of the receiving Solr node.
	Message        string                         `json:"message,omitempty"`
	Msg            string                         `json:"msg,omitempty"`
	Name           string                         `json:"name,omitempty"`
	Properties     Properties                     `json:"properties,omitempty"` // Properties and values associated with alias.
	RequestID      string                         `json:"requestid,omitempty"`
	Response       interface{}                    `json:"response,omitempty"`
	ResponseHeader ResponseHeader                 `json:"responseHeader,omitempty"`
	Schema         Properties                     `json:"schema,omitempty"`
	Similarity     interface{}                    `json:"similarity,omitempty"`
	Snapshot       string                         `json:"snapshot,omitempty"`  // The name of the snapshot to be created | deleted.
	Snapshots      map[string]SnapshotInformation `json:"snapshots,omitempty"` // The collection of snapshots found for the requested core.
	Status         string                         `json:"STATUS,omitempty"`
	Success        interface{}                    `json:"success,omitempty"`
	UniqueKey      string                         `json:"uniqueKey,omitempty"`
	Value          string                         `json:"value,omitempty"` // Property value.
	Version        float64                        `json:"version,omitempty"`
	Warning        string                         `json:"warning,omitempty"`
	ZKVersion      int                            `json:"zkversion,omitempty"`
}

type ResponseHeader struct {
	Status         int  `json:"status,omitempty"`
	QTime          int  `json:"QTime,omitempty"`
	PartialResults bool `json:"partialResults,omitempty"`
}

type SnapshotInformation struct {
	IndexDirPath     string `json:"indexDirPath,omitempty"`     // The path to the directory containing the index files.
	GenerationNumber int    `json:"generationNumber,omitempty"` // The generation value for the snapshot.
}
