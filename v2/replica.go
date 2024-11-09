package v2

type ReplicaAPI struct {
	client *Client
}

type CreateReplicaRequestBody struct {
	Name               string     `json:"name,omitempty"`
	Type               string     `json:"type,omitempty"`
	InstanceDir        string     `json:"instanceDir,omitempty"`
	DataDir            string     `json:"dataDir,omitempty"`
	ULogDir            string     `json:"ulogDir,omitempty"`
	Route              string     `json:"route,omitempty"`
	NrtReplicas        int        `json:"nrtReplicas,omitempty"`
	TLogReplicas       int        `json:"tlogReplicas,omitempty"`
	PullReplicas       int        `json:"pullReplicas,omitempty"`
	WaitForFinalState  bool       `json:"waitForFinalState,omitempty"`
	FollowAliases      bool       `json:"followAliases,omitempty"`
	Async              string     `json:"async,omitempty"`
	Node               string     `json:"node,omitempty"`
	SkipNodeAssignment bool       `json:"skipNodeAssignment,omitempty"`
	Properties         Properties `json:"properties,omitempty"`
	NodeSet            []string   `json:"nodeSet,omitempty"`
}

// Creates a new replica of an existing shard.
//
// collName - The name of the collection the replica belongs to.
//
// shardName - The name of the shard the replica belongs to.
//
// body - The value of the replica to create
//
// POST /collections/{collName}/shards/{shardName}/replicas
func (r *ReplicaAPI) Create(collName string, shardName string, body CreateReplicaRequestBody) {
}

type DeleteReplicaParams struct {
	FollowAliases     bool   `url:"followAliases,omitempty"`
	DeleteInstanceDir bool   `url:"deleteInstanceDir,omitempty"`
	DeleteDataDir     bool   `url:"deleteDataDir,omitempty"`
	DeleteIndex       bool   `url:"deleteIndex,omitempty"`
	OnlyIfDown        bool   `url:"onlyIfDown,omitempty"`
	Async             string `url:"async,omitempty"`
}

// Delete an single replica by name
//
// collName - The name of the collection the replica belongs to.
//
// shardName - The name of the shard the replica belongs to.
//
// replicaName - The replica, e.g., core_node1.
//
// DELETE /collections/{collName}/shards/{shardName}/replicas/{replicaName}
func (r *ReplicaAPI) Delete(collName string, shardName string, replicaName string, params DeleteReplicaParams) {
}

type DeleteReplicaManyParams struct {
	Count             int    `url:"count,omitempty"`
	FollowAliases     bool   `url:"followAliases,omitempty"`
	DeleteInstanceDir bool   `url:"deleteInstanceDir,omitempty"`
	DeleteDataDir     bool   `url:"deleteDataDir,omitempty"`
	DeleteIndex       bool   `url:"deleteIndex,omitempty"`
	OnlyIfDown        bool   `url:"onlyIfDown,omitempty"`
	Async             string `url:"async,omitempty"`
}

// Delete one or more replicas from the specified collection and shard
//
// collName - The name of the collection the replica belongs to.
//
// shardName - The name of the shard the replica belongs to.
//
// params - The configuration for deleting
//
// DELETE /collections/{collName}/shards/{shardName}/replicas
func (r *ReplicaAPI) DeleteMany(collName string, shardName string, params DeleteReplicaManyParams) {}

type AddPropertyRequestBody struct {
	Value       string `json:"value" validate:"required"` // The value to assign to the property.
	ShardUnique bool   `json:"shardUnique,omitempty"`     // If true, then setting this property in one replica will remove the property from all other replicas in that shard. The default is false.\nThere is one pre-defined property preferredLeader for which shardUnique is forced to true and an error returned if shardUnique is explicitly set to false.
}

// Adds a property to the specified replica.
//
// collName - The name of the collection the replica belongs to.
//
// shardName - The name of the shard the replica belongs to.
//
// replicaName - The replica, e.g., core_node1.
//
// propName - The name of the property to add.
//
// body - The value of the replica property to create or update
//
// PUT /collections/{collName}/shards/{shardName}/replicas/{replicaName}/properties/{propName}
func (r *ReplicaAPI) AddProperty(collName string, shardName string, replicaName string, propName string, body AddPropertyRequestBody) {
}

// Delete an existing replica property.
//
// collName - The name of the collection the replica belongs to.
//
// shardName - The name of the shard the replica belongs to.
//
// replicaName - The replica, e.g., core_node1.
//
// propName - The name of the property to delete.
//
// DELETE /collections/{collName}/shards/{shardName}/replicas/{replicaName}/properties/{propName}
func (r *ReplicaAPI) DeleteProperty(collName string, shardName string, replicaName string, propName string) {
}

type BalanceReplicasRequestBody struct {
	Nodes             []string `json:"nodes,omitempty"`             // The set of nodes across which replicas will be balanced. Defaults to all live data nodes.
	WaitForFinalState bool     `json:"waitForFinalState,omitempty"` // If true, the request will complete only when all affected replicas become active. If false, the API will return the status of the single action, which may be before the new replica is online and active.
	Async             string   `json:"async,omitempty"`             // Request ID to track this action which will be processed asynchronously.
}

type MigrateReplicasRequestBody struct {
	SourceNodes       []string `json:"sourceNodes" validate:"required"` // The set of nodes which all replicas will be migrated off of.
	TargetNodes       []string `json:"targetNodes,omitempty"`           // A set of nodes to migrate the replicas to. If this is not provided, then the API will use the live data nodes not in 'sourceNodes'.
	WaitForFinalState bool     `json:"waitForFinalState,omitempty"`     // If true, the request will complete only when all affected replicas become active. If false, the API will return the status of the single action, which may be before the new replicas are online and active.
	Async             string   `json:"async,omitempty"`                 // Request ID to track this action which will be processed asynchronously.
}
