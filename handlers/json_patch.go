package handlers

// JSONPatch is a format for expressing a sequence of operations to apply to a
// target JSON document. Refer to RFC 6902 for details.
type JSONPatch []JSONPatchOperation

// Append adds op as the last operation in jp.
func (jp *JSONPatch) Append(op JSONPatchOperation) {
	*jp = append(*jp, op)
}

// JSONPatchOperation is a step in a JSONPatch. Refer to RFC 6902 for details.
type JSONPatchOperation struct {
	Op    string      `json:"op"`
	Path  string      `json:"path"`
	Value interface{} `json:"value,omitempty"`
}
