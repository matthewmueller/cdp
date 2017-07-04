// Code generated by cdpgen. DO NOT EDIT.

package dom

import (
	"github.com/mafredri/cdp/protocol/runtime"
)

// GetDocumentArgs represents the arguments for GetDocument in the DOM domain.
type GetDocumentArgs struct {
	Depth  *int  `json:"depth,omitempty"`  // The maximum depth at which children should be retrieved, defaults to 1. Use -1 for the entire subtree or provide an integer larger than 0.
	Pierce *bool `json:"pierce,omitempty"` // Whether or not iframes and shadow roots should be traversed when returning the subtree (default is false).
}

// NewGetDocumentArgs initializes GetDocumentArgs with the required arguments.
func NewGetDocumentArgs() *GetDocumentArgs {
	args := new(GetDocumentArgs)

	return args
}

// SetDepth sets the Depth optional argument. The maximum depth at which children should be retrieved, defaults to 1. Use -1 for the entire subtree or provide an integer larger than 0.
func (a *GetDocumentArgs) SetDepth(depth int) *GetDocumentArgs {
	a.Depth = &depth
	return a
}

// SetPierce sets the Pierce optional argument. Whether or not iframes and shadow roots should be traversed when returning the subtree (default is false).
func (a *GetDocumentArgs) SetPierce(pierce bool) *GetDocumentArgs {
	a.Pierce = &pierce
	return a
}

// GetDocumentReply represents the return values for GetDocument in the DOM domain.
type GetDocumentReply struct {
	Root Node `json:"root"` // Resulting node.
}

// GetFlattenedDocumentArgs represents the arguments for GetFlattenedDocument in the DOM domain.
type GetFlattenedDocumentArgs struct {
	Depth  *int  `json:"depth,omitempty"`  // The maximum depth at which children should be retrieved, defaults to 1. Use -1 for the entire subtree or provide an integer larger than 0.
	Pierce *bool `json:"pierce,omitempty"` // Whether or not iframes and shadow roots should be traversed when returning the subtree (default is false).
}

// NewGetFlattenedDocumentArgs initializes GetFlattenedDocumentArgs with the required arguments.
func NewGetFlattenedDocumentArgs() *GetFlattenedDocumentArgs {
	args := new(GetFlattenedDocumentArgs)

	return args
}

// SetDepth sets the Depth optional argument. The maximum depth at which children should be retrieved, defaults to 1. Use -1 for the entire subtree or provide an integer larger than 0.
func (a *GetFlattenedDocumentArgs) SetDepth(depth int) *GetFlattenedDocumentArgs {
	a.Depth = &depth
	return a
}

// SetPierce sets the Pierce optional argument. Whether or not iframes and shadow roots should be traversed when returning the subtree (default is false).
func (a *GetFlattenedDocumentArgs) SetPierce(pierce bool) *GetFlattenedDocumentArgs {
	a.Pierce = &pierce
	return a
}

// GetFlattenedDocumentReply represents the return values for GetFlattenedDocument in the DOM domain.
type GetFlattenedDocumentReply struct {
	Nodes []Node `json:"nodes"` // Resulting node.
}

// CollectClassNamesFromSubtreeArgs represents the arguments for CollectClassNamesFromSubtree in the DOM domain.
type CollectClassNamesFromSubtreeArgs struct {
	NodeID NodeID `json:"nodeId"` // Id of the node to collect class names.
}

// NewCollectClassNamesFromSubtreeArgs initializes CollectClassNamesFromSubtreeArgs with the required arguments.
func NewCollectClassNamesFromSubtreeArgs(nodeID NodeID) *CollectClassNamesFromSubtreeArgs {
	args := new(CollectClassNamesFromSubtreeArgs)
	args.NodeID = nodeID
	return args
}

// CollectClassNamesFromSubtreeReply represents the return values for CollectClassNamesFromSubtree in the DOM domain.
type CollectClassNamesFromSubtreeReply struct {
	ClassNames []string `json:"classNames"` // Class name list.
}

// RequestChildNodesArgs represents the arguments for RequestChildNodes in the DOM domain.
type RequestChildNodesArgs struct {
	NodeID NodeID `json:"nodeId"`           // Id of the node to get children for.
	Depth  *int   `json:"depth,omitempty"`  // The maximum depth at which children should be retrieved, defaults to 1. Use -1 for the entire subtree or provide an integer larger than 0.
	Pierce *bool  `json:"pierce,omitempty"` // Whether or not iframes and shadow roots should be traversed when returning the sub-tree (default is false).
}

// NewRequestChildNodesArgs initializes RequestChildNodesArgs with the required arguments.
func NewRequestChildNodesArgs(nodeID NodeID) *RequestChildNodesArgs {
	args := new(RequestChildNodesArgs)
	args.NodeID = nodeID
	return args
}

// SetDepth sets the Depth optional argument. The maximum depth at which children should be retrieved, defaults to 1. Use -1 for the entire subtree or provide an integer larger than 0.
func (a *RequestChildNodesArgs) SetDepth(depth int) *RequestChildNodesArgs {
	a.Depth = &depth
	return a
}

// SetPierce sets the Pierce optional argument. Whether or not iframes and shadow roots should be traversed when returning the sub-tree (default is false).
func (a *RequestChildNodesArgs) SetPierce(pierce bool) *RequestChildNodesArgs {
	a.Pierce = &pierce
	return a
}

// QuerySelectorArgs represents the arguments for QuerySelector in the DOM domain.
type QuerySelectorArgs struct {
	NodeID   NodeID `json:"nodeId"`   // Id of the node to query upon.
	Selector string `json:"selector"` // Selector string.
}

// NewQuerySelectorArgs initializes QuerySelectorArgs with the required arguments.
func NewQuerySelectorArgs(nodeID NodeID, selector string) *QuerySelectorArgs {
	args := new(QuerySelectorArgs)
	args.NodeID = nodeID
	args.Selector = selector
	return args
}

// QuerySelectorReply represents the return values for QuerySelector in the DOM domain.
type QuerySelectorReply struct {
	NodeID NodeID `json:"nodeId"` // Query selector result.
}

// QuerySelectorAllArgs represents the arguments for QuerySelectorAll in the DOM domain.
type QuerySelectorAllArgs struct {
	NodeID   NodeID `json:"nodeId"`   // Id of the node to query upon.
	Selector string `json:"selector"` // Selector string.
}

// NewQuerySelectorAllArgs initializes QuerySelectorAllArgs with the required arguments.
func NewQuerySelectorAllArgs(nodeID NodeID, selector string) *QuerySelectorAllArgs {
	args := new(QuerySelectorAllArgs)
	args.NodeID = nodeID
	args.Selector = selector
	return args
}

// QuerySelectorAllReply represents the return values for QuerySelectorAll in the DOM domain.
type QuerySelectorAllReply struct {
	NodeIDs []NodeID `json:"nodeIds"` // Query selector result.
}

// SetNodeNameArgs represents the arguments for SetNodeName in the DOM domain.
type SetNodeNameArgs struct {
	NodeID NodeID `json:"nodeId"` // Id of the node to set name for.
	Name   string `json:"name"`   // New node's name.
}

// NewSetNodeNameArgs initializes SetNodeNameArgs with the required arguments.
func NewSetNodeNameArgs(nodeID NodeID, name string) *SetNodeNameArgs {
	args := new(SetNodeNameArgs)
	args.NodeID = nodeID
	args.Name = name
	return args
}

// SetNodeNameReply represents the return values for SetNodeName in the DOM domain.
type SetNodeNameReply struct {
	NodeID NodeID `json:"nodeId"` // New node's id.
}

// SetNodeValueArgs represents the arguments for SetNodeValue in the DOM domain.
type SetNodeValueArgs struct {
	NodeID NodeID `json:"nodeId"` // Id of the node to set value for.
	Value  string `json:"value"`  // New node's value.
}

// NewSetNodeValueArgs initializes SetNodeValueArgs with the required arguments.
func NewSetNodeValueArgs(nodeID NodeID, value string) *SetNodeValueArgs {
	args := new(SetNodeValueArgs)
	args.NodeID = nodeID
	args.Value = value
	return args
}

// RemoveNodeArgs represents the arguments for RemoveNode in the DOM domain.
type RemoveNodeArgs struct {
	NodeID NodeID `json:"nodeId"` // Id of the node to remove.
}

// NewRemoveNodeArgs initializes RemoveNodeArgs with the required arguments.
func NewRemoveNodeArgs(nodeID NodeID) *RemoveNodeArgs {
	args := new(RemoveNodeArgs)
	args.NodeID = nodeID
	return args
}

// SetAttributeValueArgs represents the arguments for SetAttributeValue in the DOM domain.
type SetAttributeValueArgs struct {
	NodeID NodeID `json:"nodeId"` // Id of the element to set attribute for.
	Name   string `json:"name"`   // Attribute name.
	Value  string `json:"value"`  // Attribute value.
}

// NewSetAttributeValueArgs initializes SetAttributeValueArgs with the required arguments.
func NewSetAttributeValueArgs(nodeID NodeID, name string, value string) *SetAttributeValueArgs {
	args := new(SetAttributeValueArgs)
	args.NodeID = nodeID
	args.Name = name
	args.Value = value
	return args
}

// SetAttributesAsTextArgs represents the arguments for SetAttributesAsText in the DOM domain.
type SetAttributesAsTextArgs struct {
	NodeID NodeID  `json:"nodeId"`         // Id of the element to set attributes for.
	Text   string  `json:"text"`           // Text with a number of attributes. Will parse this text using HTML parser.
	Name   *string `json:"name,omitempty"` // Attribute name to replace with new attributes derived from text in case text parsed successfully.
}

// NewSetAttributesAsTextArgs initializes SetAttributesAsTextArgs with the required arguments.
func NewSetAttributesAsTextArgs(nodeID NodeID, text string) *SetAttributesAsTextArgs {
	args := new(SetAttributesAsTextArgs)
	args.NodeID = nodeID
	args.Text = text
	return args
}

// SetName sets the Name optional argument. Attribute name to replace with new attributes derived from text in case text parsed successfully.
func (a *SetAttributesAsTextArgs) SetName(name string) *SetAttributesAsTextArgs {
	a.Name = &name
	return a
}

// RemoveAttributeArgs represents the arguments for RemoveAttribute in the DOM domain.
type RemoveAttributeArgs struct {
	NodeID NodeID `json:"nodeId"` // Id of the element to remove attribute from.
	Name   string `json:"name"`   // Name of the attribute to remove.
}

// NewRemoveAttributeArgs initializes RemoveAttributeArgs with the required arguments.
func NewRemoveAttributeArgs(nodeID NodeID, name string) *RemoveAttributeArgs {
	args := new(RemoveAttributeArgs)
	args.NodeID = nodeID
	args.Name = name
	return args
}

// GetOuterHTMLArgs represents the arguments for GetOuterHTML in the DOM domain.
type GetOuterHTMLArgs struct {
	NodeID NodeID `json:"nodeId"` // Id of the node to get markup for.
}

// NewGetOuterHTMLArgs initializes GetOuterHTMLArgs with the required arguments.
func NewGetOuterHTMLArgs(nodeID NodeID) *GetOuterHTMLArgs {
	args := new(GetOuterHTMLArgs)
	args.NodeID = nodeID
	return args
}

// GetOuterHTMLReply represents the return values for GetOuterHTML in the DOM domain.
type GetOuterHTMLReply struct {
	OuterHTML string `json:"outerHTML"` // Outer HTML markup.
}

// SetOuterHTMLArgs represents the arguments for SetOuterHTML in the DOM domain.
type SetOuterHTMLArgs struct {
	NodeID    NodeID `json:"nodeId"`    // Id of the node to set markup for.
	OuterHTML string `json:"outerHTML"` // Outer HTML markup to set.
}

// NewSetOuterHTMLArgs initializes SetOuterHTMLArgs with the required arguments.
func NewSetOuterHTMLArgs(nodeID NodeID, outerHTML string) *SetOuterHTMLArgs {
	args := new(SetOuterHTMLArgs)
	args.NodeID = nodeID
	args.OuterHTML = outerHTML
	return args
}

// PerformSearchArgs represents the arguments for PerformSearch in the DOM domain.
type PerformSearchArgs struct {
	Query                     string `json:"query"`                               // Plain text or query selector or XPath search query.
	IncludeUserAgentShadowDOM *bool  `json:"includeUserAgentShadowDOM,omitempty"` // True to search in user agent shadow DOM.
}

// NewPerformSearchArgs initializes PerformSearchArgs with the required arguments.
func NewPerformSearchArgs(query string) *PerformSearchArgs {
	args := new(PerformSearchArgs)
	args.Query = query
	return args
}

// SetIncludeUserAgentShadowDOM sets the IncludeUserAgentShadowDOM optional argument. True to search in user agent shadow DOM.
func (a *PerformSearchArgs) SetIncludeUserAgentShadowDOM(includeUserAgentShadowDOM bool) *PerformSearchArgs {
	a.IncludeUserAgentShadowDOM = &includeUserAgentShadowDOM
	return a
}

// PerformSearchReply represents the return values for PerformSearch in the DOM domain.
type PerformSearchReply struct {
	SearchID    string `json:"searchId"`    // Unique search session identifier.
	ResultCount int    `json:"resultCount"` // Number of search results.
}

// GetSearchResultsArgs represents the arguments for GetSearchResults in the DOM domain.
type GetSearchResultsArgs struct {
	SearchID  string `json:"searchId"`  // Unique search session identifier.
	FromIndex int    `json:"fromIndex"` // Start index of the search result to be returned.
	ToIndex   int    `json:"toIndex"`   // End index of the search result to be returned.
}

// NewGetSearchResultsArgs initializes GetSearchResultsArgs with the required arguments.
func NewGetSearchResultsArgs(searchID string, fromIndex int, toIndex int) *GetSearchResultsArgs {
	args := new(GetSearchResultsArgs)
	args.SearchID = searchID
	args.FromIndex = fromIndex
	args.ToIndex = toIndex
	return args
}

// GetSearchResultsReply represents the return values for GetSearchResults in the DOM domain.
type GetSearchResultsReply struct {
	NodeIDs []NodeID `json:"nodeIds"` // Ids of the search result nodes.
}

// DiscardSearchResultsArgs represents the arguments for DiscardSearchResults in the DOM domain.
type DiscardSearchResultsArgs struct {
	SearchID string `json:"searchId"` // Unique search session identifier.
}

// NewDiscardSearchResultsArgs initializes DiscardSearchResultsArgs with the required arguments.
func NewDiscardSearchResultsArgs(searchID string) *DiscardSearchResultsArgs {
	args := new(DiscardSearchResultsArgs)
	args.SearchID = searchID
	return args
}

// RequestNodeArgs represents the arguments for RequestNode in the DOM domain.
type RequestNodeArgs struct {
	ObjectID runtime.RemoteObjectID `json:"objectId"` // JavaScript object id to convert into node.
}

// NewRequestNodeArgs initializes RequestNodeArgs with the required arguments.
func NewRequestNodeArgs(objectID runtime.RemoteObjectID) *RequestNodeArgs {
	args := new(RequestNodeArgs)
	args.ObjectID = objectID
	return args
}

// RequestNodeReply represents the return values for RequestNode in the DOM domain.
type RequestNodeReply struct {
	NodeID NodeID `json:"nodeId"` // Node id for given object.
}

// PushNodeByPathToFrontendArgs represents the arguments for PushNodeByPathToFrontend in the DOM domain.
type PushNodeByPathToFrontendArgs struct {
	Path string `json:"path"` // Path to node in the proprietary format.
}

// NewPushNodeByPathToFrontendArgs initializes PushNodeByPathToFrontendArgs with the required arguments.
func NewPushNodeByPathToFrontendArgs(path string) *PushNodeByPathToFrontendArgs {
	args := new(PushNodeByPathToFrontendArgs)
	args.Path = path
	return args
}

// PushNodeByPathToFrontendReply represents the return values for PushNodeByPathToFrontend in the DOM domain.
type PushNodeByPathToFrontendReply struct {
	NodeID NodeID `json:"nodeId"` // Id of the node for given path.
}

// PushNodesByBackendIdsToFrontendArgs represents the arguments for PushNodesByBackendIdsToFrontend in the DOM domain.
type PushNodesByBackendIdsToFrontendArgs struct {
	BackendNodeIDs []BackendNodeID `json:"backendNodeIds"` // The array of backend node ids.
}

// NewPushNodesByBackendIdsToFrontendArgs initializes PushNodesByBackendIdsToFrontendArgs with the required arguments.
func NewPushNodesByBackendIdsToFrontendArgs(backendNodeIDs []BackendNodeID) *PushNodesByBackendIdsToFrontendArgs {
	args := new(PushNodesByBackendIdsToFrontendArgs)
	args.BackendNodeIDs = backendNodeIDs
	return args
}

// PushNodesByBackendIdsToFrontendReply represents the return values for PushNodesByBackendIdsToFrontend in the DOM domain.
type PushNodesByBackendIdsToFrontendReply struct {
	NodeIDs []NodeID `json:"nodeIds"` // The array of ids of pushed nodes that correspond to the backend ids specified in backendNodeIds.
}

// SetInspectedNodeArgs represents the arguments for SetInspectedNode in the DOM domain.
type SetInspectedNodeArgs struct {
	NodeID NodeID `json:"nodeId"` // DOM node id to be accessible by means of $x command line API.
}

// NewSetInspectedNodeArgs initializes SetInspectedNodeArgs with the required arguments.
func NewSetInspectedNodeArgs(nodeID NodeID) *SetInspectedNodeArgs {
	args := new(SetInspectedNodeArgs)
	args.NodeID = nodeID
	return args
}

// ResolveNodeArgs represents the arguments for ResolveNode in the DOM domain.
type ResolveNodeArgs struct {
	NodeID      NodeID  `json:"nodeId"`                // Id of the node to resolve.
	ObjectGroup *string `json:"objectGroup,omitempty"` // Symbolic group name that can be used to release multiple objects.
}

// NewResolveNodeArgs initializes ResolveNodeArgs with the required arguments.
func NewResolveNodeArgs(nodeID NodeID) *ResolveNodeArgs {
	args := new(ResolveNodeArgs)
	args.NodeID = nodeID
	return args
}

// SetObjectGroup sets the ObjectGroup optional argument. Symbolic group name that can be used to release multiple objects.
func (a *ResolveNodeArgs) SetObjectGroup(objectGroup string) *ResolveNodeArgs {
	a.ObjectGroup = &objectGroup
	return a
}

// ResolveNodeReply represents the return values for ResolveNode in the DOM domain.
type ResolveNodeReply struct {
	Object runtime.RemoteObject `json:"object"` // JavaScript object wrapper for given node.
}

// GetAttributesArgs represents the arguments for GetAttributes in the DOM domain.
type GetAttributesArgs struct {
	NodeID NodeID `json:"nodeId"` // Id of the node to retrieve attibutes for.
}

// NewGetAttributesArgs initializes GetAttributesArgs with the required arguments.
func NewGetAttributesArgs(nodeID NodeID) *GetAttributesArgs {
	args := new(GetAttributesArgs)
	args.NodeID = nodeID
	return args
}

// GetAttributesReply represents the return values for GetAttributes in the DOM domain.
type GetAttributesReply struct {
	Attributes []string `json:"attributes"` // An interleaved array of node attribute names and values.
}

// CopyToArgs represents the arguments for CopyTo in the DOM domain.
type CopyToArgs struct {
	NodeID             NodeID  `json:"nodeId"`                       // Id of the node to copy.
	TargetNodeID       NodeID  `json:"targetNodeId"`                 // Id of the element to drop the copy into.
	InsertBeforeNodeID *NodeID `json:"insertBeforeNodeId,omitempty"` // Drop the copy before this node (if absent, the copy becomes the last child of targetNodeId).
}

// NewCopyToArgs initializes CopyToArgs with the required arguments.
func NewCopyToArgs(nodeID NodeID, targetNodeID NodeID) *CopyToArgs {
	args := new(CopyToArgs)
	args.NodeID = nodeID
	args.TargetNodeID = targetNodeID
	return args
}

// SetInsertBeforeNodeID sets the InsertBeforeNodeID optional argument. Drop the copy before this node (if absent, the copy becomes the last child of targetNodeId).
func (a *CopyToArgs) SetInsertBeforeNodeID(insertBeforeNodeID NodeID) *CopyToArgs {
	a.InsertBeforeNodeID = &insertBeforeNodeID
	return a
}

// CopyToReply represents the return values for CopyTo in the DOM domain.
type CopyToReply struct {
	NodeID NodeID `json:"nodeId"` // Id of the node clone.
}

// MoveToArgs represents the arguments for MoveTo in the DOM domain.
type MoveToArgs struct {
	NodeID             NodeID  `json:"nodeId"`                       // Id of the node to move.
	TargetNodeID       NodeID  `json:"targetNodeId"`                 // Id of the element to drop the moved node into.
	InsertBeforeNodeID *NodeID `json:"insertBeforeNodeId,omitempty"` // Drop node before this one (if absent, the moved node becomes the last child of targetNodeId).
}

// NewMoveToArgs initializes MoveToArgs with the required arguments.
func NewMoveToArgs(nodeID NodeID, targetNodeID NodeID) *MoveToArgs {
	args := new(MoveToArgs)
	args.NodeID = nodeID
	args.TargetNodeID = targetNodeID
	return args
}

// SetInsertBeforeNodeID sets the InsertBeforeNodeID optional argument. Drop node before this one (if absent, the moved node becomes the last child of targetNodeId).
func (a *MoveToArgs) SetInsertBeforeNodeID(insertBeforeNodeID NodeID) *MoveToArgs {
	a.InsertBeforeNodeID = &insertBeforeNodeID
	return a
}

// MoveToReply represents the return values for MoveTo in the DOM domain.
type MoveToReply struct {
	NodeID NodeID `json:"nodeId"` // New id of the moved node.
}

// FocusArgs represents the arguments for Focus in the DOM domain.
type FocusArgs struct {
	NodeID NodeID `json:"nodeId"` // Id of the node to focus.
}

// NewFocusArgs initializes FocusArgs with the required arguments.
func NewFocusArgs(nodeID NodeID) *FocusArgs {
	args := new(FocusArgs)
	args.NodeID = nodeID
	return args
}

// SetFileInputFilesArgs represents the arguments for SetFileInputFiles in the DOM domain.
type SetFileInputFilesArgs struct {
	NodeID NodeID   `json:"nodeId"` // Id of the file input node to set files for.
	Files  []string `json:"files"`  // Array of file paths to set.
}

// NewSetFileInputFilesArgs initializes SetFileInputFilesArgs with the required arguments.
func NewSetFileInputFilesArgs(nodeID NodeID, files []string) *SetFileInputFilesArgs {
	args := new(SetFileInputFilesArgs)
	args.NodeID = nodeID
	args.Files = files
	return args
}

// GetBoxModelArgs represents the arguments for GetBoxModel in the DOM domain.
type GetBoxModelArgs struct {
	NodeID NodeID `json:"nodeId"` // Id of the node to get box model for.
}

// NewGetBoxModelArgs initializes GetBoxModelArgs with the required arguments.
func NewGetBoxModelArgs(nodeID NodeID) *GetBoxModelArgs {
	args := new(GetBoxModelArgs)
	args.NodeID = nodeID
	return args
}

// GetBoxModelReply represents the return values for GetBoxModel in the DOM domain.
type GetBoxModelReply struct {
	Model BoxModel `json:"model"` // Box model for the node.
}

// GetNodeForLocationArgs represents the arguments for GetNodeForLocation in the DOM domain.
type GetNodeForLocationArgs struct {
	X                         int   `json:"x"`                                   // X coordinate.
	Y                         int   `json:"y"`                                   // Y coordinate.
	IncludeUserAgentShadowDOM *bool `json:"includeUserAgentShadowDOM,omitempty"` // False to skip to the nearest non-UA shadow root ancestor (default: false).
}

// NewGetNodeForLocationArgs initializes GetNodeForLocationArgs with the required arguments.
func NewGetNodeForLocationArgs(x int, y int) *GetNodeForLocationArgs {
	args := new(GetNodeForLocationArgs)
	args.X = x
	args.Y = y
	return args
}

// SetIncludeUserAgentShadowDOM sets the IncludeUserAgentShadowDOM optional argument. False to skip to the nearest non-UA shadow root ancestor (default: false).
func (a *GetNodeForLocationArgs) SetIncludeUserAgentShadowDOM(includeUserAgentShadowDOM bool) *GetNodeForLocationArgs {
	a.IncludeUserAgentShadowDOM = &includeUserAgentShadowDOM
	return a
}

// GetNodeForLocationReply represents the return values for GetNodeForLocation in the DOM domain.
type GetNodeForLocationReply struct {
	NodeID NodeID `json:"nodeId"` // Id of the node at given coordinates.
}

// GetRelayoutBoundaryArgs represents the arguments for GetRelayoutBoundary in the DOM domain.
type GetRelayoutBoundaryArgs struct {
	NodeID NodeID `json:"nodeId"` // Id of the node.
}

// NewGetRelayoutBoundaryArgs initializes GetRelayoutBoundaryArgs with the required arguments.
func NewGetRelayoutBoundaryArgs(nodeID NodeID) *GetRelayoutBoundaryArgs {
	args := new(GetRelayoutBoundaryArgs)
	args.NodeID = nodeID
	return args
}

// GetRelayoutBoundaryReply represents the return values for GetRelayoutBoundary in the DOM domain.
type GetRelayoutBoundaryReply struct {
	NodeID NodeID `json:"nodeId"` // Relayout boundary node id for the given node.
}
