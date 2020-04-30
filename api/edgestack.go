package portainer

// EdgeStackRelatedEndpoints returns a list of endpoints related to this Edge stack
func EdgeStackRelatedEndpoints(edgeStack EdgeStack, endpoints []Endpoint, endpointGroups []EndpointGroup, edgeGroups []EdgeGroup) ([]EndpointID, error) {
	edgeStackEndpoints := []EndpointID{}

	for _, edgeGroupID := range edgeStack.EdgeGroups {
		var edgeGroup *EdgeGroup

		for _, group := range edgeGroups {
			if group.ID == edgeGroupID {
				edgeGroup = &group
				break
			}
		}

		if edgeGroup == nil {
			return nil, Error("Edge group was not found")
		}

		edgeStackEndpoints = append(edgeStackEndpoints, EdgeGroupRelatedEndpoints(*edgeGroup, endpoints, endpointGroups)...)
	}

	return edgeStackEndpoints, nil
}
