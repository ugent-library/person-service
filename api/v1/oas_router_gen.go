// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/ogen-go/ogen/uri"
)

// ServeHTTP serves http request as defined by OpenAPI v3 specification,
// calling handler that matches the path or returning not found error.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	elem := r.URL.Path
	elemIsEscaped := false
	if rawPath := r.URL.RawPath; rawPath != "" {
		if normalized, ok := uri.NormalizeEscapedPath(rawPath); ok {
			elem = normalized
			elemIsEscaped = strings.ContainsRune(elem, '%')
		}
	}
	if prefix := s.cfg.Prefix; len(prefix) > 0 {
		if strings.HasPrefix(elem, prefix) {
			// Cut prefix from the path.
			elem = strings.TrimPrefix(elem, prefix)
		} else {
			// Prefix doesn't match.
			s.notFound(w, r)
			return
		}
	}
	if len(elem) == 0 {
		s.notFound(w, r)
		return
	}

	// Static code generated router with unwrapped path search.
	switch {
	default:
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/"
			if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				break
			}
			switch elem[0] {
			case 'a': // Prefix: "add-"
				if l := len("add-"); len(elem) >= l && elem[0:l] == "add-" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'o': // Prefix: "organization"
					if l := len("organization"); len(elem) >= l && elem[0:l] == "organization" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "POST":
							s.handleAddOrganizationRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "POST")
						}

						return
					}
				case 'p': // Prefix: "person"
					if l := len("person"); len(elem) >= l && elem[0:l] == "person" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "POST":
							s.handleAddPersonRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "POST")
						}

						return
					}
				}
			case 'g': // Prefix: "get-"
				if l := len("get-"); len(elem) >= l && elem[0:l] == "get-" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'o': // Prefix: "organization"
					if l := len("organization"); len(elem) >= l && elem[0:l] == "organization" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch r.Method {
						case "POST":
							s.handleGetOrganizationRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "POST")
						}

						return
					}
					switch elem[0] {
					case 's': // Prefix: "s"
						if l := len("s"); len(elem) >= l && elem[0:l] == "s" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch r.Method {
							case "POST":
								s.handleGetOrganizationsRequest([0]string{}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "POST")
							}

							return
						}
						switch elem[0] {
						case '-': // Prefix: "-by-id"
							if l := len("-by-id"); len(elem) >= l && elem[0:l] == "-by-id" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "POST":
									s.handleGetOrganizationsByIdRequest([0]string{}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, "POST")
								}

								return
							}
						}
					}
				case 'p': // Prefix: "pe"
					if l := len("pe"); len(elem) >= l && elem[0:l] == "pe" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'o': // Prefix: "ople"
						if l := len("ople"); len(elem) >= l && elem[0:l] == "ople" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch r.Method {
							case "POST":
								s.handleGetPeopleRequest([0]string{}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "POST")
							}

							return
						}
						switch elem[0] {
						case '-': // Prefix: "-by-id"
							if l := len("-by-id"); len(elem) >= l && elem[0:l] == "-by-id" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "POST":
									s.handleGetPeopleByIdRequest([0]string{}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, "POST")
								}

								return
							}
						}
					case 'r': // Prefix: "rson"
						if l := len("rson"); len(elem) >= l && elem[0:l] == "rson" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "POST":
								s.handleGetPersonRequest([0]string{}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "POST")
							}

							return
						}
					}
				}
			case 's': // Prefix: "s"
				if l := len("s"); len(elem) >= l && elem[0:l] == "s" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'e': // Prefix: "et-person-"
					if l := len("et-person-"); len(elem) >= l && elem[0:l] == "et-person-" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'o': // Prefix: "orcid"
						if l := len("orcid"); len(elem) >= l && elem[0:l] == "orcid" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch r.Method {
							case "POST":
								s.handleSetPersonOrcidRequest([0]string{}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "POST")
							}

							return
						}
						switch elem[0] {
						case '-': // Prefix: "-token"
							if l := len("-token"); len(elem) >= l && elem[0:l] == "-token" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "POST":
									s.handleSetPersonOrcidTokenRequest([0]string{}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, "POST")
								}

								return
							}
						}
					case 'r': // Prefix: "role"
						if l := len("role"); len(elem) >= l && elem[0:l] == "role" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "POST":
								s.handleSetPersonRoleRequest([0]string{}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "POST")
							}

							return
						}
					case 's': // Prefix: "settings"
						if l := len("settings"); len(elem) >= l && elem[0:l] == "settings" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "POST":
								s.handleSetPersonSettingsRequest([0]string{}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "POST")
							}

							return
						}
					}
				case 'u': // Prefix: "uggest-"
					if l := len("uggest-"); len(elem) >= l && elem[0:l] == "uggest-" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'o': // Prefix: "organizations"
						if l := len("organizations"); len(elem) >= l && elem[0:l] == "organizations" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "POST":
								s.handleSuggestOrganizationsRequest([0]string{}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "POST")
							}

							return
						}
					case 'p': // Prefix: "people"
						if l := len("people"); len(elem) >= l && elem[0:l] == "people" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "POST":
								s.handleSuggestPeopleRequest([0]string{}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "POST")
							}

							return
						}
					}
				}
			}
		}
	}
	s.notFound(w, r)
}

// Route is route object.
type Route struct {
	name        string
	operationID string
	pathPattern string
	count       int
	args        [0]string
}

// Name returns ogen operation name.
//
// It is guaranteed to be unique and not empty.
func (r Route) Name() string {
	return r.name
}

// OperationID returns OpenAPI operationId.
func (r Route) OperationID() string {
	return r.operationID
}

// PathPattern returns OpenAPI path.
func (r Route) PathPattern() string {
	return r.pathPattern
}

// Args returns parsed arguments.
func (r Route) Args() []string {
	return r.args[:r.count]
}

// FindRoute finds Route for given method and path.
//
// Note: this method does not unescape path or handle reserved characters in path properly. Use FindPath instead.
func (s *Server) FindRoute(method, path string) (Route, bool) {
	return s.FindPath(method, &url.URL{Path: path})
}

// FindPath finds Route for given method and URL.
func (s *Server) FindPath(method string, u *url.URL) (r Route, _ bool) {
	var (
		elem = u.Path
		args = r.args
	)
	if rawPath := u.RawPath; rawPath != "" {
		if normalized, ok := uri.NormalizeEscapedPath(rawPath); ok {
			elem = normalized
		}
		defer func() {
			for i, arg := range r.args[:r.count] {
				if unescaped, err := url.PathUnescape(arg); err == nil {
					r.args[i] = unescaped
				}
			}
		}()
	}

	// Static code generated router with unwrapped path search.
	switch {
	default:
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/"
			if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				break
			}
			switch elem[0] {
			case 'a': // Prefix: "add-"
				if l := len("add-"); len(elem) >= l && elem[0:l] == "add-" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'o': // Prefix: "organization"
					if l := len("organization"); len(elem) >= l && elem[0:l] == "organization" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "POST":
							// Leaf: AddOrganization
							r.name = "AddOrganization"
							r.operationID = "AddOrganization"
							r.pathPattern = "/add-organization"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
				case 'p': // Prefix: "person"
					if l := len("person"); len(elem) >= l && elem[0:l] == "person" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "POST":
							// Leaf: AddPerson
							r.name = "AddPerson"
							r.operationID = "AddPerson"
							r.pathPattern = "/add-person"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
				}
			case 'g': // Prefix: "get-"
				if l := len("get-"); len(elem) >= l && elem[0:l] == "get-" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'o': // Prefix: "organization"
					if l := len("organization"); len(elem) >= l && elem[0:l] == "organization" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "POST":
							r.name = "GetOrganization"
							r.operationID = "GetOrganization"
							r.pathPattern = "/get-organization"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
					switch elem[0] {
					case 's': // Prefix: "s"
						if l := len("s"); len(elem) >= l && elem[0:l] == "s" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch method {
							case "POST":
								r.name = "GetOrganizations"
								r.operationID = "GetOrganizations"
								r.pathPattern = "/get-organizations"
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}
						switch elem[0] {
						case '-': // Prefix: "-by-id"
							if l := len("-by-id"); len(elem) >= l && elem[0:l] == "-by-id" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								switch method {
								case "POST":
									// Leaf: GetOrganizationsById
									r.name = "GetOrganizationsById"
									r.operationID = "GetOrganizationsById"
									r.pathPattern = "/get-organizations-by-id"
									r.args = args
									r.count = 0
									return r, true
								default:
									return
								}
							}
						}
					}
				case 'p': // Prefix: "pe"
					if l := len("pe"); len(elem) >= l && elem[0:l] == "pe" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'o': // Prefix: "ople"
						if l := len("ople"); len(elem) >= l && elem[0:l] == "ople" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch method {
							case "POST":
								r.name = "GetPeople"
								r.operationID = "GetPeople"
								r.pathPattern = "/get-people"
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}
						switch elem[0] {
						case '-': // Prefix: "-by-id"
							if l := len("-by-id"); len(elem) >= l && elem[0:l] == "-by-id" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								switch method {
								case "POST":
									// Leaf: GetPeopleById
									r.name = "GetPeopleById"
									r.operationID = "GetPeopleById"
									r.pathPattern = "/get-people-by-id"
									r.args = args
									r.count = 0
									return r, true
								default:
									return
								}
							}
						}
					case 'r': // Prefix: "rson"
						if l := len("rson"); len(elem) >= l && elem[0:l] == "rson" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch method {
							case "POST":
								// Leaf: GetPerson
								r.name = "GetPerson"
								r.operationID = "GetPerson"
								r.pathPattern = "/get-person"
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}
					}
				}
			case 's': // Prefix: "s"
				if l := len("s"); len(elem) >= l && elem[0:l] == "s" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'e': // Prefix: "et-person-"
					if l := len("et-person-"); len(elem) >= l && elem[0:l] == "et-person-" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'o': // Prefix: "orcid"
						if l := len("orcid"); len(elem) >= l && elem[0:l] == "orcid" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch method {
							case "POST":
								r.name = "SetPersonOrcid"
								r.operationID = "SetPersonOrcid"
								r.pathPattern = "/set-person-orcid"
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}
						switch elem[0] {
						case '-': // Prefix: "-token"
							if l := len("-token"); len(elem) >= l && elem[0:l] == "-token" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								switch method {
								case "POST":
									// Leaf: SetPersonOrcidToken
									r.name = "SetPersonOrcidToken"
									r.operationID = "SetPersonOrcidToken"
									r.pathPattern = "/set-person-orcid-token"
									r.args = args
									r.count = 0
									return r, true
								default:
									return
								}
							}
						}
					case 'r': // Prefix: "role"
						if l := len("role"); len(elem) >= l && elem[0:l] == "role" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch method {
							case "POST":
								// Leaf: SetPersonRole
								r.name = "SetPersonRole"
								r.operationID = "SetPersonRole"
								r.pathPattern = "/set-person-role"
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}
					case 's': // Prefix: "settings"
						if l := len("settings"); len(elem) >= l && elem[0:l] == "settings" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch method {
							case "POST":
								// Leaf: SetPersonSettings
								r.name = "SetPersonSettings"
								r.operationID = "SetPersonSettings"
								r.pathPattern = "/set-person-settings"
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}
					}
				case 'u': // Prefix: "uggest-"
					if l := len("uggest-"); len(elem) >= l && elem[0:l] == "uggest-" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'o': // Prefix: "organizations"
						if l := len("organizations"); len(elem) >= l && elem[0:l] == "organizations" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch method {
							case "POST":
								// Leaf: SuggestOrganizations
								r.name = "SuggestOrganizations"
								r.operationID = "SuggestOrganizations"
								r.pathPattern = "/suggest-organizations"
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}
					case 'p': // Prefix: "people"
						if l := len("people"); len(elem) >= l && elem[0:l] == "people" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch method {
							case "POST":
								// Leaf: SuggestPeople
								r.name = "SuggestPeople"
								r.operationID = "SuggestPeople"
								r.pathPattern = "/suggest-people"
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}
					}
				}
			}
		}
	}
	return r, false
}
