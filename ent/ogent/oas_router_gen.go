// Code generated by ogen, DO NOT EDIT.

package ogent

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"math/big"
	"math/bits"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"github.com/google/uuid"
	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/otelogen"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

// No-op definition for keeping imports.
var (
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = sort.Ints
	_ = http.MethodGet
	_ = io.Copy
	_ = json.Marshal
	_ = bytes.NewReader
	_ = strconv.ParseInt
	_ = time.Time{}
	_ = conv.ToInt32
	_ = uuid.UUID{}
	_ = uri.PathEncoder{}
	_ = url.URL{}
	_ = math.Mod
	_ = bits.LeadingZeros64
	_ = big.Rat{}
	_ = validate.Int{}
	_ = ht.NewRequest
	_ = net.IP{}
	_ = otelogen.Version
	_ = attribute.KeyValue{}
	_ = trace.TraceIDFromHex
	_ = otel.GetTracerProvider
	_ = metric.NewNoopMeterProvider
	_ = regexp.MustCompile
	_ = jx.Null
	_ = sync.Pool{}
	_ = codes.Unset
)

func (s *Server) notFound(w http.ResponseWriter, r *http.Request) {
	s.cfg.NotFound(w, r)
}

// ServeHTTP serves http request as defined by OpenAPI v3 specification,
// calling handler that matches the path or returning not found error.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	elem := r.URL.Path
	if len(elem) == 0 {
		s.notFound(w, r)
		return
	}
	args := [1]string{}
	// Static code generated router with unwrapped path search.
	switch r.Method {
	case "DELETE":
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
				s.handleDeleteCarRequest([1]string{
					args[0],
				}, w, r)

				return
			}
			switch elem[0] {
			case 'b': // Prefix: "battles/"
				if l := len("battles/"); len(elem) >= l && elem[0:l] == "battles/" {
					elem = elem[l:]
				} else {
					break
				}

				// Param: "id"
				// Leaf parameter
				args[0] = elem
				elem = ""

				if len(elem) == 0 {
					// Leaf: DeleteBattle
					s.handleDeleteBattleRequest([1]string{
						args[0],
					}, w, r)

					return
				}
			case 'c': // Prefix: "cars/"
				if l := len("cars/"); len(elem) >= l && elem[0:l] == "cars/" {
					elem = elem[l:]
				} else {
					break
				}

				// Param: "id"
				// Leaf parameter
				args[0] = elem
				elem = ""

				if len(elem) == 0 {
					// Leaf: DeleteCar
					s.handleDeleteCarRequest([1]string{
						args[0],
					}, w, r)

					return
				}
			case 'g': // Prefix: "groups/"
				if l := len("groups/"); len(elem) >= l && elem[0:l] == "groups/" {
					elem = elem[l:]
				} else {
					break
				}

				// Param: "id"
				// Leaf parameter
				args[0] = elem
				elem = ""

				if len(elem) == 0 {
					// Leaf: DeleteGroup
					s.handleDeleteGroupRequest([1]string{
						args[0],
					}, w, r)

					return
				}
			case 'p': // Prefix: "p"
				if l := len("p"); len(elem) >= l && elem[0:l] == "p" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					s.handleDeletePokemonRequest([1]string{
						args[0],
					}, w, r)

					return
				}
				switch elem[0] {
				case 'e': // Prefix: "ets/"
					if l := len("ets/"); len(elem) >= l && elem[0:l] == "ets/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "id"
					// Leaf parameter
					args[0] = elem
					elem = ""

					if len(elem) == 0 {
						// Leaf: DeletePet
						s.handleDeletePetRequest([1]string{
							args[0],
						}, w, r)

						return
					}
				case 'o': // Prefix: "okemons/"
					if l := len("okemons/"); len(elem) >= l && elem[0:l] == "okemons/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "id"
					// Leaf parameter
					args[0] = elem
					elem = ""

					if len(elem) == 0 {
						// Leaf: DeletePokemon
						s.handleDeletePokemonRequest([1]string{
							args[0],
						}, w, r)

						return
					}
				}
			case 'u': // Prefix: "users/"
				if l := len("users/"); len(elem) >= l && elem[0:l] == "users/" {
					elem = elem[l:]
				} else {
					break
				}

				// Param: "id"
				// Leaf parameter
				args[0] = elem
				elem = ""

				if len(elem) == 0 {
					// Leaf: DeleteUser
					s.handleDeleteUserRequest([1]string{
						args[0],
					}, w, r)

					return
				}
			}
		}
	case "GET":
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
				s.handleListBattleRequest([0]string{}, w, r)

				return
			}
			switch elem[0] {
			case 'b': // Prefix: "battles"
				if l := len("battles"); len(elem) >= l && elem[0:l] == "battles" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					s.handleListBattleRequest([0]string{}, w, r)

					return
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "id"
					// Match until "/"
					idx := strings.IndexByte(elem, '/')
					if idx < 0 {
						idx = len(elem)
					}
					args[0] = elem[:idx]
					elem = elem[idx:]

					if len(elem) == 0 {
						s.handleReadBattleRequest([1]string{
							args[0],
						}, w, r)

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
							s.handleReadBattleOponentRequest([1]string{
								args[0],
							}, w, r)

							return
						}
						switch elem[0] {
						case 'c': // Prefix: "contender"
							if l := len("contender"); len(elem) >= l && elem[0:l] == "contender" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf: ReadBattleContender
								s.handleReadBattleContenderRequest([1]string{
									args[0],
								}, w, r)

								return
							}
						case 'o': // Prefix: "oponent"
							if l := len("oponent"); len(elem) >= l && elem[0:l] == "oponent" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf: ReadBattleOponent
								s.handleReadBattleOponentRequest([1]string{
									args[0],
								}, w, r)

								return
							}
						}
					}
				}
			case 'c': // Prefix: "cars"
				if l := len("cars"); len(elem) >= l && elem[0:l] == "cars" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					s.handleListCarRequest([0]string{}, w, r)

					return
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "id"
					// Match until "/"
					idx := strings.IndexByte(elem, '/')
					if idx < 0 {
						idx = len(elem)
					}
					args[0] = elem[:idx]
					elem = elem[idx:]

					if len(elem) == 0 {
						s.handleReadCarRequest([1]string{
							args[0],
						}, w, r)

						return
					}
					switch elem[0] {
					case '/': // Prefix: "/cars"
						if l := len("/cars"); len(elem) >= l && elem[0:l] == "/cars" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf: ListCarCars
							s.handleListCarCarsRequest([1]string{
								args[0],
							}, w, r)

							return
						}
					}
				}
			case 'd': // Prefix: "db-health"
				if l := len("db-health"); len(elem) >= l && elem[0:l] == "db-health" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf: DBHealth
					s.handleDBHealthRequest([0]string{}, w, r)

					return
				}
			case 'g': // Prefix: "groups"
				if l := len("groups"); len(elem) >= l && elem[0:l] == "groups" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					s.handleListGroupRequest([0]string{}, w, r)

					return
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "id"
					// Match until "/"
					idx := strings.IndexByte(elem, '/')
					if idx < 0 {
						idx = len(elem)
					}
					args[0] = elem[:idx]
					elem = elem[idx:]

					if len(elem) == 0 {
						s.handleReadGroupRequest([1]string{
							args[0],
						}, w, r)

						return
					}
					switch elem[0] {
					case '/': // Prefix: "/users"
						if l := len("/users"); len(elem) >= l && elem[0:l] == "/users" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf: ListGroupUsers
							s.handleListGroupUsersRequest([1]string{
								args[0],
							}, w, r)

							return
						}
					}
				}
			case 'p': // Prefix: "p"
				if l := len("p"); len(elem) >= l && elem[0:l] == "p" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					s.handleListPokemonRequest([0]string{}, w, r)

					return
				}
				switch elem[0] {
				case 'e': // Prefix: "ets"
					if l := len("ets"); len(elem) >= l && elem[0:l] == "ets" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						s.handleListPetRequest([0]string{}, w, r)

						return
					}
					switch elem[0] {
					case '/': // Prefix: "/"
						if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
							elem = elem[l:]
						} else {
							break
						}

						// Param: "id"
						// Leaf parameter
						args[0] = elem
						elem = ""

						if len(elem) == 0 {
							// Leaf: ReadPet
							s.handleReadPetRequest([1]string{
								args[0],
							}, w, r)

							return
						}
					}
				case 'o': // Prefix: "okemons"
					if l := len("okemons"); len(elem) >= l && elem[0:l] == "okemons" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						s.handleListPokemonRequest([0]string{}, w, r)

						return
					}
					switch elem[0] {
					case '/': // Prefix: "/"
						if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
							elem = elem[l:]
						} else {
							break
						}

						// Param: "id"
						// Match until "/"
						idx := strings.IndexByte(elem, '/')
						if idx < 0 {
							idx = len(elem)
						}
						args[0] = elem[:idx]
						elem = elem[idx:]

						if len(elem) == 0 {
							s.handleReadPokemonRequest([1]string{
								args[0],
							}, w, r)

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
								s.handleListPokemonOpponentsRequest([1]string{
									args[0],
								}, w, r)

								return
							}
							switch elem[0] {
							case 'f': // Prefix: "fights"
								if l := len("fights"); len(elem) >= l && elem[0:l] == "fights" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf: ListPokemonFights
									s.handleListPokemonFightsRequest([1]string{
										args[0],
									}, w, r)

									return
								}
							case 'o': // Prefix: "opponents"
								if l := len("opponents"); len(elem) >= l && elem[0:l] == "opponents" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf: ListPokemonOpponents
									s.handleListPokemonOpponentsRequest([1]string{
										args[0],
									}, w, r)

									return
								}
							}
						}
					}
				}
			case 'u': // Prefix: "users"
				if l := len("users"); len(elem) >= l && elem[0:l] == "users" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					s.handleListUserRequest([0]string{}, w, r)

					return
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "id"
					// Match until "/"
					idx := strings.IndexByte(elem, '/')
					if idx < 0 {
						idx = len(elem)
					}
					args[0] = elem[:idx]
					elem = elem[idx:]

					if len(elem) == 0 {
						s.handleReadUserRequest([1]string{
							args[0],
						}, w, r)

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
							s.handleListUserGroupsRequest([1]string{
								args[0],
							}, w, r)

							return
						}
						switch elem[0] {
						case 'c': // Prefix: "cars"
							if l := len("cars"); len(elem) >= l && elem[0:l] == "cars" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf: ListUserCars
								s.handleListUserCarsRequest([1]string{
									args[0],
								}, w, r)

								return
							}
						case 'g': // Prefix: "groups"
							if l := len("groups"); len(elem) >= l && elem[0:l] == "groups" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf: ListUserGroups
								s.handleListUserGroupsRequest([1]string{
									args[0],
								}, w, r)

								return
							}
						}
					}
				}
			}
		}
	case "PATCH":
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
				s.handleUpdateCarRequest([1]string{
					args[0],
				}, w, r)

				return
			}
			switch elem[0] {
			case 'b': // Prefix: "battles/"
				if l := len("battles/"); len(elem) >= l && elem[0:l] == "battles/" {
					elem = elem[l:]
				} else {
					break
				}

				// Param: "id"
				// Leaf parameter
				args[0] = elem
				elem = ""

				if len(elem) == 0 {
					// Leaf: UpdateBattle
					s.handleUpdateBattleRequest([1]string{
						args[0],
					}, w, r)

					return
				}
			case 'c': // Prefix: "cars/"
				if l := len("cars/"); len(elem) >= l && elem[0:l] == "cars/" {
					elem = elem[l:]
				} else {
					break
				}

				// Param: "id"
				// Leaf parameter
				args[0] = elem
				elem = ""

				if len(elem) == 0 {
					// Leaf: UpdateCar
					s.handleUpdateCarRequest([1]string{
						args[0],
					}, w, r)

					return
				}
			case 'g': // Prefix: "groups/"
				if l := len("groups/"); len(elem) >= l && elem[0:l] == "groups/" {
					elem = elem[l:]
				} else {
					break
				}

				// Param: "id"
				// Leaf parameter
				args[0] = elem
				elem = ""

				if len(elem) == 0 {
					// Leaf: UpdateGroup
					s.handleUpdateGroupRequest([1]string{
						args[0],
					}, w, r)

					return
				}
			case 'p': // Prefix: "p"
				if l := len("p"); len(elem) >= l && elem[0:l] == "p" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					s.handleUpdatePokemonRequest([1]string{
						args[0],
					}, w, r)

					return
				}
				switch elem[0] {
				case 'e': // Prefix: "ets/"
					if l := len("ets/"); len(elem) >= l && elem[0:l] == "ets/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "id"
					// Leaf parameter
					args[0] = elem
					elem = ""

					if len(elem) == 0 {
						// Leaf: UpdatePet
						s.handleUpdatePetRequest([1]string{
							args[0],
						}, w, r)

						return
					}
				case 'o': // Prefix: "okemons/"
					if l := len("okemons/"); len(elem) >= l && elem[0:l] == "okemons/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "id"
					// Leaf parameter
					args[0] = elem
					elem = ""

					if len(elem) == 0 {
						// Leaf: UpdatePokemon
						s.handleUpdatePokemonRequest([1]string{
							args[0],
						}, w, r)

						return
					}
				}
			case 'u': // Prefix: "users/"
				if l := len("users/"); len(elem) >= l && elem[0:l] == "users/" {
					elem = elem[l:]
				} else {
					break
				}

				// Param: "id"
				// Leaf parameter
				args[0] = elem
				elem = ""

				if len(elem) == 0 {
					// Leaf: UpdateUser
					s.handleUpdateUserRequest([1]string{
						args[0],
					}, w, r)

					return
				}
			}
		}
	case "POST":
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
				s.handleCreateCarRequest([0]string{}, w, r)

				return
			}
			switch elem[0] {
			case 'b': // Prefix: "battles"
				if l := len("battles"); len(elem) >= l && elem[0:l] == "battles" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf: CreateBattle
					s.handleCreateBattleRequest([0]string{}, w, r)

					return
				}
			case 'c': // Prefix: "cars"
				if l := len("cars"); len(elem) >= l && elem[0:l] == "cars" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf: CreateCar
					s.handleCreateCarRequest([0]string{}, w, r)

					return
				}
			case 'g': // Prefix: "groups"
				if l := len("groups"); len(elem) >= l && elem[0:l] == "groups" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf: CreateGroup
					s.handleCreateGroupRequest([0]string{}, w, r)

					return
				}
			case 'p': // Prefix: "p"
				if l := len("p"); len(elem) >= l && elem[0:l] == "p" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					s.handleCreatePokemonRequest([0]string{}, w, r)

					return
				}
				switch elem[0] {
				case 'e': // Prefix: "ets"
					if l := len("ets"); len(elem) >= l && elem[0:l] == "ets" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf: CreatePet
						s.handleCreatePetRequest([0]string{}, w, r)

						return
					}
				case 'o': // Prefix: "okemons"
					if l := len("okemons"); len(elem) >= l && elem[0:l] == "okemons" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf: CreatePokemon
						s.handleCreatePokemonRequest([0]string{}, w, r)

						return
					}
				}
			case 'u': // Prefix: "users"
				if l := len("users"); len(elem) >= l && elem[0:l] == "users" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf: CreateUser
					s.handleCreateUserRequest([0]string{}, w, r)

					return
				}
			}
		}
	}
	s.notFound(w, r)
}

// Route is route object.
type Route struct {
	name  string
	count int
	args  [1]string
}

// OperationID returns OpenAPI operationId.
func (r Route) OperationID() string {
	return r.name
}

// Args returns parsed arguments.
func (r Route) Args() []string {
	return r.args[:r.count]
}

// FindRoute finds Route for given method and path.
func (s *Server) FindRoute(method, path string) (r Route, _ bool) {
	var (
		args = [1]string{}
		elem = path
	)
	r.args = args

	// Static code generated router with unwrapped path search.
	switch method {
	case "DELETE":
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
				r.name = "DeleteCar"
				r.args = args
				r.count = 0
				return r, true
			}
			switch elem[0] {
			case 'b': // Prefix: "battles/"
				if l := len("battles/"); len(elem) >= l && elem[0:l] == "battles/" {
					elem = elem[l:]
				} else {
					break
				}

				// Param: "id"
				// Leaf parameter
				args[0] = elem
				elem = ""

				if len(elem) == 0 {
					// Leaf: DeleteBattle
					r.name = "DeleteBattle"
					r.args = args
					r.count = 1
					return r, true
				}
			case 'c': // Prefix: "cars/"
				if l := len("cars/"); len(elem) >= l && elem[0:l] == "cars/" {
					elem = elem[l:]
				} else {
					break
				}

				// Param: "id"
				// Leaf parameter
				args[0] = elem
				elem = ""

				if len(elem) == 0 {
					// Leaf: DeleteCar
					r.name = "DeleteCar"
					r.args = args
					r.count = 1
					return r, true
				}
			case 'g': // Prefix: "groups/"
				if l := len("groups/"); len(elem) >= l && elem[0:l] == "groups/" {
					elem = elem[l:]
				} else {
					break
				}

				// Param: "id"
				// Leaf parameter
				args[0] = elem
				elem = ""

				if len(elem) == 0 {
					// Leaf: DeleteGroup
					r.name = "DeleteGroup"
					r.args = args
					r.count = 1
					return r, true
				}
			case 'p': // Prefix: "p"
				if l := len("p"); len(elem) >= l && elem[0:l] == "p" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					r.name = "DeletePokemon"
					r.args = args
					r.count = 0
					return r, true
				}
				switch elem[0] {
				case 'e': // Prefix: "ets/"
					if l := len("ets/"); len(elem) >= l && elem[0:l] == "ets/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "id"
					// Leaf parameter
					args[0] = elem
					elem = ""

					if len(elem) == 0 {
						// Leaf: DeletePet
						r.name = "DeletePet"
						r.args = args
						r.count = 1
						return r, true
					}
				case 'o': // Prefix: "okemons/"
					if l := len("okemons/"); len(elem) >= l && elem[0:l] == "okemons/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "id"
					// Leaf parameter
					args[0] = elem
					elem = ""

					if len(elem) == 0 {
						// Leaf: DeletePokemon
						r.name = "DeletePokemon"
						r.args = args
						r.count = 1
						return r, true
					}
				}
			case 'u': // Prefix: "users/"
				if l := len("users/"); len(elem) >= l && elem[0:l] == "users/" {
					elem = elem[l:]
				} else {
					break
				}

				// Param: "id"
				// Leaf parameter
				args[0] = elem
				elem = ""

				if len(elem) == 0 {
					// Leaf: DeleteUser
					r.name = "DeleteUser"
					r.args = args
					r.count = 1
					return r, true
				}
			}
		}
	case "GET":
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
				r.name = "ListBattle"
				r.args = args
				r.count = 0
				return r, true
			}
			switch elem[0] {
			case 'b': // Prefix: "battles"
				if l := len("battles"); len(elem) >= l && elem[0:l] == "battles" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					r.name = "ListBattle"
					r.args = args
					r.count = 0
					return r, true
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "id"
					// Match until "/"
					idx := strings.IndexByte(elem, '/')
					if idx < 0 {
						idx = len(elem)
					}
					args[0] = elem[:idx]
					elem = elem[idx:]

					if len(elem) == 0 {
						r.name = "ReadBattle"
						r.args = args
						r.count = 1
						return r, true
					}
					switch elem[0] {
					case '/': // Prefix: "/"
						if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							r.name = "ReadBattleOponent"
							r.args = args
							r.count = 1
							return r, true
						}
						switch elem[0] {
						case 'c': // Prefix: "contender"
							if l := len("contender"); len(elem) >= l && elem[0:l] == "contender" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf: ReadBattleContender
								r.name = "ReadBattleContender"
								r.args = args
								r.count = 1
								return r, true
							}
						case 'o': // Prefix: "oponent"
							if l := len("oponent"); len(elem) >= l && elem[0:l] == "oponent" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf: ReadBattleOponent
								r.name = "ReadBattleOponent"
								r.args = args
								r.count = 1
								return r, true
							}
						}
					}
				}
			case 'c': // Prefix: "cars"
				if l := len("cars"); len(elem) >= l && elem[0:l] == "cars" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					r.name = "ListCar"
					r.args = args
					r.count = 0
					return r, true
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "id"
					// Match until "/"
					idx := strings.IndexByte(elem, '/')
					if idx < 0 {
						idx = len(elem)
					}
					args[0] = elem[:idx]
					elem = elem[idx:]

					if len(elem) == 0 {
						r.name = "ReadCar"
						r.args = args
						r.count = 1
						return r, true
					}
					switch elem[0] {
					case '/': // Prefix: "/cars"
						if l := len("/cars"); len(elem) >= l && elem[0:l] == "/cars" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf: ListCarCars
							r.name = "ListCarCars"
							r.args = args
							r.count = 1
							return r, true
						}
					}
				}
			case 'd': // Prefix: "db-health"
				if l := len("db-health"); len(elem) >= l && elem[0:l] == "db-health" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf: DBHealth
					r.name = "DBHealth"
					r.args = args
					r.count = 0
					return r, true
				}
			case 'g': // Prefix: "groups"
				if l := len("groups"); len(elem) >= l && elem[0:l] == "groups" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					r.name = "ListGroup"
					r.args = args
					r.count = 0
					return r, true
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "id"
					// Match until "/"
					idx := strings.IndexByte(elem, '/')
					if idx < 0 {
						idx = len(elem)
					}
					args[0] = elem[:idx]
					elem = elem[idx:]

					if len(elem) == 0 {
						r.name = "ReadGroup"
						r.args = args
						r.count = 1
						return r, true
					}
					switch elem[0] {
					case '/': // Prefix: "/users"
						if l := len("/users"); len(elem) >= l && elem[0:l] == "/users" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf: ListGroupUsers
							r.name = "ListGroupUsers"
							r.args = args
							r.count = 1
							return r, true
						}
					}
				}
			case 'p': // Prefix: "p"
				if l := len("p"); len(elem) >= l && elem[0:l] == "p" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					r.name = "ListPokemon"
					r.args = args
					r.count = 0
					return r, true
				}
				switch elem[0] {
				case 'e': // Prefix: "ets"
					if l := len("ets"); len(elem) >= l && elem[0:l] == "ets" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						r.name = "ListPet"
						r.args = args
						r.count = 0
						return r, true
					}
					switch elem[0] {
					case '/': // Prefix: "/"
						if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
							elem = elem[l:]
						} else {
							break
						}

						// Param: "id"
						// Leaf parameter
						args[0] = elem
						elem = ""

						if len(elem) == 0 {
							// Leaf: ReadPet
							r.name = "ReadPet"
							r.args = args
							r.count = 1
							return r, true
						}
					}
				case 'o': // Prefix: "okemons"
					if l := len("okemons"); len(elem) >= l && elem[0:l] == "okemons" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						r.name = "ListPokemon"
						r.args = args
						r.count = 0
						return r, true
					}
					switch elem[0] {
					case '/': // Prefix: "/"
						if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
							elem = elem[l:]
						} else {
							break
						}

						// Param: "id"
						// Match until "/"
						idx := strings.IndexByte(elem, '/')
						if idx < 0 {
							idx = len(elem)
						}
						args[0] = elem[:idx]
						elem = elem[idx:]

						if len(elem) == 0 {
							r.name = "ReadPokemon"
							r.args = args
							r.count = 1
							return r, true
						}
						switch elem[0] {
						case '/': // Prefix: "/"
							if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								r.name = "ListPokemonOpponents"
								r.args = args
								r.count = 1
								return r, true
							}
							switch elem[0] {
							case 'f': // Prefix: "fights"
								if l := len("fights"); len(elem) >= l && elem[0:l] == "fights" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf: ListPokemonFights
									r.name = "ListPokemonFights"
									r.args = args
									r.count = 1
									return r, true
								}
							case 'o': // Prefix: "opponents"
								if l := len("opponents"); len(elem) >= l && elem[0:l] == "opponents" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf: ListPokemonOpponents
									r.name = "ListPokemonOpponents"
									r.args = args
									r.count = 1
									return r, true
								}
							}
						}
					}
				}
			case 'u': // Prefix: "users"
				if l := len("users"); len(elem) >= l && elem[0:l] == "users" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					r.name = "ListUser"
					r.args = args
					r.count = 0
					return r, true
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "id"
					// Match until "/"
					idx := strings.IndexByte(elem, '/')
					if idx < 0 {
						idx = len(elem)
					}
					args[0] = elem[:idx]
					elem = elem[idx:]

					if len(elem) == 0 {
						r.name = "ReadUser"
						r.args = args
						r.count = 1
						return r, true
					}
					switch elem[0] {
					case '/': // Prefix: "/"
						if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							r.name = "ListUserGroups"
							r.args = args
							r.count = 1
							return r, true
						}
						switch elem[0] {
						case 'c': // Prefix: "cars"
							if l := len("cars"); len(elem) >= l && elem[0:l] == "cars" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf: ListUserCars
								r.name = "ListUserCars"
								r.args = args
								r.count = 1
								return r, true
							}
						case 'g': // Prefix: "groups"
							if l := len("groups"); len(elem) >= l && elem[0:l] == "groups" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf: ListUserGroups
								r.name = "ListUserGroups"
								r.args = args
								r.count = 1
								return r, true
							}
						}
					}
				}
			}
		}
	case "PATCH":
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
				r.name = "UpdateCar"
				r.args = args
				r.count = 0
				return r, true
			}
			switch elem[0] {
			case 'b': // Prefix: "battles/"
				if l := len("battles/"); len(elem) >= l && elem[0:l] == "battles/" {
					elem = elem[l:]
				} else {
					break
				}

				// Param: "id"
				// Leaf parameter
				args[0] = elem
				elem = ""

				if len(elem) == 0 {
					// Leaf: UpdateBattle
					r.name = "UpdateBattle"
					r.args = args
					r.count = 1
					return r, true
				}
			case 'c': // Prefix: "cars/"
				if l := len("cars/"); len(elem) >= l && elem[0:l] == "cars/" {
					elem = elem[l:]
				} else {
					break
				}

				// Param: "id"
				// Leaf parameter
				args[0] = elem
				elem = ""

				if len(elem) == 0 {
					// Leaf: UpdateCar
					r.name = "UpdateCar"
					r.args = args
					r.count = 1
					return r, true
				}
			case 'g': // Prefix: "groups/"
				if l := len("groups/"); len(elem) >= l && elem[0:l] == "groups/" {
					elem = elem[l:]
				} else {
					break
				}

				// Param: "id"
				// Leaf parameter
				args[0] = elem
				elem = ""

				if len(elem) == 0 {
					// Leaf: UpdateGroup
					r.name = "UpdateGroup"
					r.args = args
					r.count = 1
					return r, true
				}
			case 'p': // Prefix: "p"
				if l := len("p"); len(elem) >= l && elem[0:l] == "p" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					r.name = "UpdatePokemon"
					r.args = args
					r.count = 0
					return r, true
				}
				switch elem[0] {
				case 'e': // Prefix: "ets/"
					if l := len("ets/"); len(elem) >= l && elem[0:l] == "ets/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "id"
					// Leaf parameter
					args[0] = elem
					elem = ""

					if len(elem) == 0 {
						// Leaf: UpdatePet
						r.name = "UpdatePet"
						r.args = args
						r.count = 1
						return r, true
					}
				case 'o': // Prefix: "okemons/"
					if l := len("okemons/"); len(elem) >= l && elem[0:l] == "okemons/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "id"
					// Leaf parameter
					args[0] = elem
					elem = ""

					if len(elem) == 0 {
						// Leaf: UpdatePokemon
						r.name = "UpdatePokemon"
						r.args = args
						r.count = 1
						return r, true
					}
				}
			case 'u': // Prefix: "users/"
				if l := len("users/"); len(elem) >= l && elem[0:l] == "users/" {
					elem = elem[l:]
				} else {
					break
				}

				// Param: "id"
				// Leaf parameter
				args[0] = elem
				elem = ""

				if len(elem) == 0 {
					// Leaf: UpdateUser
					r.name = "UpdateUser"
					r.args = args
					r.count = 1
					return r, true
				}
			}
		}
	case "POST":
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
				r.name = "CreateCar"
				r.args = args
				r.count = 0
				return r, true
			}
			switch elem[0] {
			case 'b': // Prefix: "battles"
				if l := len("battles"); len(elem) >= l && elem[0:l] == "battles" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf: CreateBattle
					r.name = "CreateBattle"
					r.args = args
					r.count = 0
					return r, true
				}
			case 'c': // Prefix: "cars"
				if l := len("cars"); len(elem) >= l && elem[0:l] == "cars" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf: CreateCar
					r.name = "CreateCar"
					r.args = args
					r.count = 0
					return r, true
				}
			case 'g': // Prefix: "groups"
				if l := len("groups"); len(elem) >= l && elem[0:l] == "groups" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf: CreateGroup
					r.name = "CreateGroup"
					r.args = args
					r.count = 0
					return r, true
				}
			case 'p': // Prefix: "p"
				if l := len("p"); len(elem) >= l && elem[0:l] == "p" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					r.name = "CreatePokemon"
					r.args = args
					r.count = 0
					return r, true
				}
				switch elem[0] {
				case 'e': // Prefix: "ets"
					if l := len("ets"); len(elem) >= l && elem[0:l] == "ets" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf: CreatePet
						r.name = "CreatePet"
						r.args = args
						r.count = 0
						return r, true
					}
				case 'o': // Prefix: "okemons"
					if l := len("okemons"); len(elem) >= l && elem[0:l] == "okemons" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf: CreatePokemon
						r.name = "CreatePokemon"
						r.args = args
						r.count = 0
						return r, true
					}
				}
			case 'u': // Prefix: "users"
				if l := len("users"); len(elem) >= l && elem[0:l] == "users" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf: CreateUser
					r.name = "CreateUser"
					r.args = args
					r.count = 0
					return r, true
				}
			}
		}
	}
	return r, false
}
