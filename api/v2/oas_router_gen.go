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
	args := [1]string{}

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
			case 'o': // Prefix: "organizations"
				if l := len("organizations"); len(elem) >= l && elem[0:l] == "organizations" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch r.Method {
					case "GET":
						s.handleGetOrganizationsRequest([0]string{}, elemIsEscaped, w, r)
					default:
						s.notAllowed(w, r, "GET")
					}

					return
				}
				switch elem[0] {
				case '-': // Prefix: "-suggest"
					if l := len("-suggest"); len(elem) >= l && elem[0:l] == "-suggest" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "GET":
							s.handleSuggestOrganizationsRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}
				case '/': // Prefix: "/"
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "organizationId"
					// Leaf parameter
					args[0] = elem
					elem = ""

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "GET":
							s.handleGetOrganizationRequest([1]string{
								args[0],
							}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

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
					switch r.Method {
					case "GET":
						s.handleGetPeopleRequest([0]string{}, elemIsEscaped, w, r)
					default:
						s.notAllowed(w, r, "GET")
					}

					return
				}
				switch elem[0] {
				case '-': // Prefix: "-suggest"
					if l := len("-suggest"); len(elem) >= l && elem[0:l] == "-suggest" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "GET":
							s.handleSuggestPeopleRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}
				case '/': // Prefix: "/"
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "personId"
					// Match until "/"
					idx := strings.IndexByte(elem, '/')
					if idx < 0 {
						idx = len(elem)
					}
					args[0] = elem[:idx]
					elem = elem[idx:]

					if len(elem) == 0 {
						switch r.Method {
						case "GET":
							s.handleGetPersonRequest([1]string{
								args[0],
							}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
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
						case 'o': // Prefix: "orcid"
							if l := len("orcid"); len(elem) >= l && elem[0:l] == "orcid" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								switch r.Method {
								case "PUT":
									s.handleSetPersonOrcidRequest([1]string{
										args[0],
									}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, "PUT")
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
									case "PUT":
										s.handleSetPersonOrcidTokenRequest([1]string{
											args[0],
										}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, "PUT")
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
								case "PUT":
									s.handleSetPersonRoleRequest([1]string{
										args[0],
									}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, "PUT")
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
								case "PUT":
									s.handleSetPersonSettingsRequest([1]string{
										args[0],
									}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, "PUT")
								}

								return
							}
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
	args        [1]string
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
			case 'o': // Prefix: "organizations"
				if l := len("organizations"); len(elem) >= l && elem[0:l] == "organizations" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch method {
					case "GET":
						r.name = "GetOrganizations"
						r.operationID = "getOrganizations"
						r.pathPattern = "/organizations"
						r.args = args
						r.count = 0
						return r, true
					default:
						return
					}
				}
				switch elem[0] {
				case '-': // Prefix: "-suggest"
					if l := len("-suggest"); len(elem) >= l && elem[0:l] == "-suggest" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "GET":
							// Leaf: SuggestOrganizations
							r.name = "SuggestOrganizations"
							r.operationID = "suggestOrganizations"
							r.pathPattern = "/organizations-suggest"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
				case '/': // Prefix: "/"
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "organizationId"
					// Leaf parameter
					args[0] = elem
					elem = ""

					if len(elem) == 0 {
						switch method {
						case "GET":
							// Leaf: GetOrganization
							r.name = "GetOrganization"
							r.operationID = "getOrganization"
							r.pathPattern = "/organizations/{organizationId}"
							r.args = args
							r.count = 1
							return r, true
						default:
							return
						}
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
					case "GET":
						r.name = "GetPeople"
						r.operationID = "getPeople"
						r.pathPattern = "/people"
						r.args = args
						r.count = 0
						return r, true
					default:
						return
					}
				}
				switch elem[0] {
				case '-': // Prefix: "-suggest"
					if l := len("-suggest"); len(elem) >= l && elem[0:l] == "-suggest" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "GET":
							// Leaf: SuggestPeople
							r.name = "SuggestPeople"
							r.operationID = "suggestPeople"
							r.pathPattern = "/people-suggest"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
				case '/': // Prefix: "/"
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "personId"
					// Match until "/"
					idx := strings.IndexByte(elem, '/')
					if idx < 0 {
						idx = len(elem)
					}
					args[0] = elem[:idx]
					elem = elem[idx:]

					if len(elem) == 0 {
						switch method {
						case "GET":
							r.name = "GetPerson"
							r.operationID = "getPerson"
							r.pathPattern = "/people/{personId}"
							r.args = args
							r.count = 1
							return r, true
						default:
							return
						}
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
						case 'o': // Prefix: "orcid"
							if l := len("orcid"); len(elem) >= l && elem[0:l] == "orcid" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								switch method {
								case "PUT":
									r.name = "SetPersonOrcid"
									r.operationID = "setPersonOrcid"
									r.pathPattern = "/people/{personId}/orcid"
									r.args = args
									r.count = 1
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
									case "PUT":
										// Leaf: SetPersonOrcidToken
										r.name = "SetPersonOrcidToken"
										r.operationID = "setPersonOrcidToken"
										r.pathPattern = "/people/{personId}/orcid-token"
										r.args = args
										r.count = 1
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
								case "PUT":
									// Leaf: SetPersonRole
									r.name = "SetPersonRole"
									r.operationID = "setPersonRole"
									r.pathPattern = "/people/{personId}/role"
									r.args = args
									r.count = 1
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
								case "PUT":
									// Leaf: SetPersonSettings
									r.name = "SetPersonSettings"
									r.operationID = "setPersonSettings"
									r.pathPattern = "/people/{personId}/settings"
									r.args = args
									r.count = 1
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
	}
	return r, false
}
