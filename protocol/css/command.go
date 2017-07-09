// Code generated by cdpgen. DO NOT EDIT.

package css

import (
	"github.com/mafredri/cdp/protocol/dom"
)

// GetMatchedStylesForNodeArgs represents the arguments for GetMatchedStylesForNode in the CSS domain.
type GetMatchedStylesForNodeArgs struct {
	NodeID dom.NodeID `json:"nodeId"` //
}

// NewGetMatchedStylesForNodeArgs initializes GetMatchedStylesForNodeArgs with the required arguments.
func NewGetMatchedStylesForNodeArgs(nodeID dom.NodeID) *GetMatchedStylesForNodeArgs {
	args := new(GetMatchedStylesForNodeArgs)
	args.NodeID = nodeID
	return args
}

// GetMatchedStylesForNodeReply represents the return values for GetMatchedStylesForNode in the CSS domain.
type GetMatchedStylesForNodeReply struct {
	InlineStyle       *Style                 `json:"inlineStyle,omitempty"`       // Inline style for the specified DOM node.
	AttributesStyle   *Style                 `json:"attributesStyle,omitempty"`   // Attribute-defined element style (e.g. resulting from "width=20 height=100%").
	MatchedCSSRules   []RuleMatch            `json:"matchedCSSRules,omitempty"`   // CSS rules matching this node, from all applicable stylesheets.
	PseudoElements    []PseudoElementMatches `json:"pseudoElements,omitempty"`    // Pseudo style matches for this node.
	Inherited         []InheritedStyleEntry  `json:"inherited,omitempty"`         // A chain of inherited styles (from the immediate node parent up to the DOM tree root).
	CSSKeyframesRules []KeyframesRule        `json:"cssKeyframesRules,omitempty"` // A list of CSS keyframed animations matching this node.
}

// GetInlineStylesForNodeArgs represents the arguments for GetInlineStylesForNode in the CSS domain.
type GetInlineStylesForNodeArgs struct {
	NodeID dom.NodeID `json:"nodeId"` //
}

// NewGetInlineStylesForNodeArgs initializes GetInlineStylesForNodeArgs with the required arguments.
func NewGetInlineStylesForNodeArgs(nodeID dom.NodeID) *GetInlineStylesForNodeArgs {
	args := new(GetInlineStylesForNodeArgs)
	args.NodeID = nodeID
	return args
}

// GetInlineStylesForNodeReply represents the return values for GetInlineStylesForNode in the CSS domain.
type GetInlineStylesForNodeReply struct {
	InlineStyle     *Style `json:"inlineStyle,omitempty"`     // Inline style for the specified DOM node.
	AttributesStyle *Style `json:"attributesStyle,omitempty"` // Attribute-defined element style (e.g. resulting from "width=20 height=100%").
}

// GetComputedStyleForNodeArgs represents the arguments for GetComputedStyleForNode in the CSS domain.
type GetComputedStyleForNodeArgs struct {
	NodeID dom.NodeID `json:"nodeId"` //
}

// NewGetComputedStyleForNodeArgs initializes GetComputedStyleForNodeArgs with the required arguments.
func NewGetComputedStyleForNodeArgs(nodeID dom.NodeID) *GetComputedStyleForNodeArgs {
	args := new(GetComputedStyleForNodeArgs)
	args.NodeID = nodeID
	return args
}

// GetComputedStyleForNodeReply represents the return values for GetComputedStyleForNode in the CSS domain.
type GetComputedStyleForNodeReply struct {
	ComputedStyle []ComputedStyleProperty `json:"computedStyle"` // Computed style for the specified DOM node.
}

// GetPlatformFontsForNodeArgs represents the arguments for GetPlatformFontsForNode in the CSS domain.
type GetPlatformFontsForNodeArgs struct {
	NodeID dom.NodeID `json:"nodeId"` //
}

// NewGetPlatformFontsForNodeArgs initializes GetPlatformFontsForNodeArgs with the required arguments.
func NewGetPlatformFontsForNodeArgs(nodeID dom.NodeID) *GetPlatformFontsForNodeArgs {
	args := new(GetPlatformFontsForNodeArgs)
	args.NodeID = nodeID
	return args
}

// GetPlatformFontsForNodeReply represents the return values for GetPlatformFontsForNode in the CSS domain.
type GetPlatformFontsForNodeReply struct {
	Fonts []PlatformFontUsage `json:"fonts"` // Usage statistics for every employed platform font.
}

// GetStyleSheetTextArgs represents the arguments for GetStyleSheetText in the CSS domain.
type GetStyleSheetTextArgs struct {
	StyleSheetID StyleSheetID `json:"styleSheetId"` //
}

// NewGetStyleSheetTextArgs initializes GetStyleSheetTextArgs with the required arguments.
func NewGetStyleSheetTextArgs(styleSheetID StyleSheetID) *GetStyleSheetTextArgs {
	args := new(GetStyleSheetTextArgs)
	args.StyleSheetID = styleSheetID
	return args
}

// GetStyleSheetTextReply represents the return values for GetStyleSheetText in the CSS domain.
type GetStyleSheetTextReply struct {
	Text string `json:"text"` // The stylesheet text.
}

// CollectClassNamesArgs represents the arguments for CollectClassNames in the CSS domain.
type CollectClassNamesArgs struct {
	StyleSheetID StyleSheetID `json:"styleSheetId"` //
}

// NewCollectClassNamesArgs initializes CollectClassNamesArgs with the required arguments.
func NewCollectClassNamesArgs(styleSheetID StyleSheetID) *CollectClassNamesArgs {
	args := new(CollectClassNamesArgs)
	args.StyleSheetID = styleSheetID
	return args
}

// CollectClassNamesReply represents the return values for CollectClassNames in the CSS domain.
type CollectClassNamesReply struct {
	ClassNames []string `json:"classNames"` // Class name list.
}

// SetStyleSheetTextArgs represents the arguments for SetStyleSheetText in the CSS domain.
type SetStyleSheetTextArgs struct {
	StyleSheetID StyleSheetID `json:"styleSheetId"` //
	Text         string       `json:"text"`         //
}

// NewSetStyleSheetTextArgs initializes SetStyleSheetTextArgs with the required arguments.
func NewSetStyleSheetTextArgs(styleSheetID StyleSheetID, text string) *SetStyleSheetTextArgs {
	args := new(SetStyleSheetTextArgs)
	args.StyleSheetID = styleSheetID
	args.Text = text
	return args
}

// SetStyleSheetTextReply represents the return values for SetStyleSheetText in the CSS domain.
type SetStyleSheetTextReply struct {
	SourceMapURL *string `json:"sourceMapURL,omitempty"` // URL of source map associated with script (if any).
}

// SetRuleSelectorArgs represents the arguments for SetRuleSelector in the CSS domain.
type SetRuleSelectorArgs struct {
	StyleSheetID StyleSheetID `json:"styleSheetId"` //
	Range        SourceRange  `json:"range"`        //
	Selector     string       `json:"selector"`     //
}

// NewSetRuleSelectorArgs initializes SetRuleSelectorArgs with the required arguments.
func NewSetRuleSelectorArgs(styleSheetID StyleSheetID, rang SourceRange, selector string) *SetRuleSelectorArgs {
	args := new(SetRuleSelectorArgs)
	args.StyleSheetID = styleSheetID
	args.Range = rang
	args.Selector = selector
	return args
}

// SetRuleSelectorReply represents the return values for SetRuleSelector in the CSS domain.
type SetRuleSelectorReply struct {
	SelectorList SelectorList `json:"selectorList"` // The resulting selector list after modification.
}

// SetKeyframeKeyArgs represents the arguments for SetKeyframeKey in the CSS domain.
type SetKeyframeKeyArgs struct {
	StyleSheetID StyleSheetID `json:"styleSheetId"` //
	Range        SourceRange  `json:"range"`        //
	KeyText      string       `json:"keyText"`      //
}

// NewSetKeyframeKeyArgs initializes SetKeyframeKeyArgs with the required arguments.
func NewSetKeyframeKeyArgs(styleSheetID StyleSheetID, rang SourceRange, keyText string) *SetKeyframeKeyArgs {
	args := new(SetKeyframeKeyArgs)
	args.StyleSheetID = styleSheetID
	args.Range = rang
	args.KeyText = keyText
	return args
}

// SetKeyframeKeyReply represents the return values for SetKeyframeKey in the CSS domain.
type SetKeyframeKeyReply struct {
	KeyText Value `json:"keyText"` // The resulting key text after modification.
}

// SetStyleTextsArgs represents the arguments for SetStyleTexts in the CSS domain.
type SetStyleTextsArgs struct {
	Edits []StyleDeclarationEdit `json:"edits"` //
}

// NewSetStyleTextsArgs initializes SetStyleTextsArgs with the required arguments.
func NewSetStyleTextsArgs(edits []StyleDeclarationEdit) *SetStyleTextsArgs {
	args := new(SetStyleTextsArgs)
	args.Edits = edits
	return args
}

// SetStyleTextsReply represents the return values for SetStyleTexts in the CSS domain.
type SetStyleTextsReply struct {
	Styles []Style `json:"styles"` // The resulting styles after modification.
}

// SetMediaTextArgs represents the arguments for SetMediaText in the CSS domain.
type SetMediaTextArgs struct {
	StyleSheetID StyleSheetID `json:"styleSheetId"` //
	Range        SourceRange  `json:"range"`        //
	Text         string       `json:"text"`         //
}

// NewSetMediaTextArgs initializes SetMediaTextArgs with the required arguments.
func NewSetMediaTextArgs(styleSheetID StyleSheetID, rang SourceRange, text string) *SetMediaTextArgs {
	args := new(SetMediaTextArgs)
	args.StyleSheetID = styleSheetID
	args.Range = rang
	args.Text = text
	return args
}

// SetMediaTextReply represents the return values for SetMediaText in the CSS domain.
type SetMediaTextReply struct {
	Media Media `json:"media"` // The resulting CSS media rule after modification.
}

// CreateStyleSheetReply represents the return values for CreateStyleSheet in the CSS domain.
type CreateStyleSheetReply struct {
	StyleSheetID StyleSheetID `json:"styleSheetId"` // Identifier of the created "via-inspector" stylesheet.
}

// AddRuleArgs represents the arguments for AddRule in the CSS domain.
type AddRuleArgs struct {
	StyleSheetID StyleSheetID `json:"styleSheetId"` // The css style sheet identifier where a new rule should be inserted.
	RuleText     string       `json:"ruleText"`     // The text of a new rule.
	Location     SourceRange  `json:"location"`     // Text position of a new rule in the target style sheet.
}

// NewAddRuleArgs initializes AddRuleArgs with the required arguments.
func NewAddRuleArgs(styleSheetID StyleSheetID, ruleText string, location SourceRange) *AddRuleArgs {
	args := new(AddRuleArgs)
	args.StyleSheetID = styleSheetID
	args.RuleText = ruleText
	args.Location = location
	return args
}

// AddRuleReply represents the return values for AddRule in the CSS domain.
type AddRuleReply struct {
	Rule Rule `json:"rule"` // The newly created rule.
}

// ForcePseudoStateArgs represents the arguments for ForcePseudoState in the CSS domain.
type ForcePseudoStateArgs struct {
	NodeID              dom.NodeID `json:"nodeId"`              // The element id for which to force the pseudo state.
	ForcedPseudoClasses []string   `json:"forcedPseudoClasses"` // Element pseudo classes to force when computing the element's style.
}

// NewForcePseudoStateArgs initializes ForcePseudoStateArgs with the required arguments.
func NewForcePseudoStateArgs(nodeID dom.NodeID, forcedPseudoClasses []string) *ForcePseudoStateArgs {
	args := new(ForcePseudoStateArgs)
	args.NodeID = nodeID
	args.ForcedPseudoClasses = forcedPseudoClasses
	return args
}

// GetMediaQueriesReply represents the return values for GetMediaQueries in the CSS domain.
type GetMediaQueriesReply struct {
	Medias []Media `json:"medias"` //
}

// SetEffectivePropertyValueForNodeArgs represents the arguments for SetEffectivePropertyValueForNode in the CSS domain.
type SetEffectivePropertyValueForNodeArgs struct {
	NodeID       dom.NodeID `json:"nodeId"`       // The element id for which to set property.
	PropertyName string     `json:"propertyName"` //
	Value        string     `json:"value"`        //
}

// NewSetEffectivePropertyValueForNodeArgs initializes SetEffectivePropertyValueForNodeArgs with the required arguments.
func NewSetEffectivePropertyValueForNodeArgs(nodeID dom.NodeID, propertyName string, value string) *SetEffectivePropertyValueForNodeArgs {
	args := new(SetEffectivePropertyValueForNodeArgs)
	args.NodeID = nodeID
	args.PropertyName = propertyName
	args.Value = value
	return args
}

// GetBackgroundColorsArgs represents the arguments for GetBackgroundColors in the CSS domain.
type GetBackgroundColorsArgs struct {
	NodeID dom.NodeID `json:"nodeId"` // Id of the node to get background colors for.
}

// NewGetBackgroundColorsArgs initializes GetBackgroundColorsArgs with the required arguments.
func NewGetBackgroundColorsArgs(nodeID dom.NodeID) *GetBackgroundColorsArgs {
	args := new(GetBackgroundColorsArgs)
	args.NodeID = nodeID
	return args
}

// GetBackgroundColorsReply represents the return values for GetBackgroundColors in the CSS domain.
type GetBackgroundColorsReply struct {
	BackgroundColors []string `json:"backgroundColors,omitempty"` // The range of background colors behind this element, if it contains any visible text. If no visible text is present, this will be undefined. In the case of a flat background color, this will consist of simply that color. In the case of a gradient, this will consist of each of the color stops. For anything more complicated, this will be an empty array. Images will be ignored (as if the image had failed to load).
}

// TakeCoverageDeltaReply represents the return values for TakeCoverageDelta in the CSS domain.
type TakeCoverageDeltaReply struct {
	Coverage []RuleUsage `json:"coverage"` //
}

// StopRuleUsageTrackingReply represents the return values for StopRuleUsageTracking in the CSS domain.
type StopRuleUsageTrackingReply struct {
	RuleUsage []RuleUsage `json:"ruleUsage"` //
}
