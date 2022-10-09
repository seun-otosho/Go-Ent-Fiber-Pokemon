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

// Ref: #/components/schemas/Battle_ContenderRead
type BattleContenderRead struct {
	ID          int       "json:\"id\""
	Name        string    "json:\"name\""
	Description string    "json:\"description\""
	Weight      float64   "json:\"weight\""
	Height      float64   "json:\"height\""
	CreatedAt   time.Time "json:\"created_at\""
	UpdatedAt   time.Time "json:\"updated_at\""
}

func (*BattleContenderRead) readBattleContenderRes() {}

// Ref: #/components/schemas/BattleCreate
type BattleCreate struct {
	ID        int       "json:\"id\""
	Result    string    "json:\"result\""
	CreatedAt time.Time "json:\"created_at\""
	UpdatedAt time.Time "json:\"updated_at\""
}

func (*BattleCreate) createBattleRes() {}

// Ref: #/components/schemas/BattleList
type BattleList struct {
	ID        int       "json:\"id\""
	Result    string    "json:\"result\""
	CreatedAt time.Time "json:\"created_at\""
	UpdatedAt time.Time "json:\"updated_at\""
}

// Ref: #/components/schemas/Battle_OponentRead
type BattleOponentRead struct {
	ID          int       "json:\"id\""
	Name        string    "json:\"name\""
	Description string    "json:\"description\""
	Weight      float64   "json:\"weight\""
	Height      float64   "json:\"height\""
	CreatedAt   time.Time "json:\"created_at\""
	UpdatedAt   time.Time "json:\"updated_at\""
}

func (*BattleOponentRead) readBattleOponentRes() {}

// Ref: #/components/schemas/BattleRead
type BattleRead struct {
	ID        int       "json:\"id\""
	Result    string    "json:\"result\""
	CreatedAt time.Time "json:\"created_at\""
	UpdatedAt time.Time "json:\"updated_at\""
}

func (*BattleRead) readBattleRes() {}

// Ref: #/components/schemas/BattleUpdate
type BattleUpdate struct {
	ID        int       "json:\"id\""
	Result    string    "json:\"result\""
	CreatedAt time.Time "json:\"created_at\""
	UpdatedAt time.Time "json:\"updated_at\""
}

func (*BattleUpdate) updateBattleRes() {}

// Ref: #/components/schemas/Car_CarsList
type CarCarsList struct {
	ID           int       "json:\"id\""
	Model        string    "json:\"model\""
	RegisteredAt time.Time "json:\"registered_at\""
}

// Ref: #/components/schemas/CarCreate
type CarCreate struct {
	ID           int       "json:\"id\""
	Model        string    "json:\"model\""
	RegisteredAt time.Time "json:\"registered_at\""
}

func (*CarCreate) createCarRes() {}

// Ref: #/components/schemas/CarList
type CarList struct {
	ID           int       "json:\"id\""
	Model        string    "json:\"model\""
	RegisteredAt time.Time "json:\"registered_at\""
}

// Ref: #/components/schemas/CarRead
type CarRead struct {
	ID           int       "json:\"id\""
	Model        string    "json:\"model\""
	RegisteredAt time.Time "json:\"registered_at\""
}

func (*CarRead) readCarRes() {}

// Ref: #/components/schemas/CarUpdate
type CarUpdate struct {
	ID           int       "json:\"id\""
	Model        string    "json:\"model\""
	RegisteredAt time.Time "json:\"registered_at\""
}

func (*CarUpdate) updateCarRes() {}

type CreateBattleReq struct {
	Result    string    "json:\"result\""
	CreatedAt time.Time "json:\"created_at\""
	UpdatedAt time.Time "json:\"updated_at\""
	Contender OptInt    "json:\"contender\""
	Oponent   OptInt    "json:\"oponent\""
}

type CreateCarReq struct {
	Model        string    "json:\"model\""
	RegisteredAt time.Time "json:\"registered_at\""
	Cars         []int     "json:\"cars\""
}

type CreateGroupReq struct {
	Name  string "json:\"name\""
	Users []int  "json:\"users\""
}

type CreatePetReq struct {
	Name string "json:\"name\""
	Age  int    "json:\"age\""
}

type CreatePokemonReq struct {
	Name        string    "json:\"name\""
	Description string    "json:\"description\""
	Weight      float64   "json:\"weight\""
	Height      float64   "json:\"height\""
	CreatedAt   time.Time "json:\"created_at\""
	UpdatedAt   time.Time "json:\"updated_at\""
	Fights      []int     "json:\"fights\""
	Opponents   []int     "json:\"opponents\""
}

type CreateUserReq struct {
	Age    int    "json:\"age\""
	Name   string "json:\"name\""
	Cars   []int  "json:\"cars\""
	Groups []int  "json:\"groups\""
}

// DBHealthNoContent is response for DBHealth operation.
type DBHealthNoContent struct{}

func (*DBHealthNoContent) dBHealthRes() {}

// DBHealthServiceUnavailable is response for DBHealth operation.
type DBHealthServiceUnavailable struct{}

func (*DBHealthServiceUnavailable) dBHealthRes() {}

// DeleteBattleNoContent is response for DeleteBattle operation.
type DeleteBattleNoContent struct{}

func (*DeleteBattleNoContent) deleteBattleRes() {}

// DeleteCarNoContent is response for DeleteCar operation.
type DeleteCarNoContent struct{}

func (*DeleteCarNoContent) deleteCarRes() {}

// DeleteGroupNoContent is response for DeleteGroup operation.
type DeleteGroupNoContent struct{}

func (*DeleteGroupNoContent) deleteGroupRes() {}

// DeletePetNoContent is response for DeletePet operation.
type DeletePetNoContent struct{}

func (*DeletePetNoContent) deletePetRes() {}

// DeletePokemonNoContent is response for DeletePokemon operation.
type DeletePokemonNoContent struct{}

func (*DeletePokemonNoContent) deletePokemonRes() {}

// DeleteUserNoContent is response for DeleteUser operation.
type DeleteUserNoContent struct{}

func (*DeleteUserNoContent) deleteUserRes() {}

// Ref: #/components/schemas/GroupCreate
type GroupCreate struct {
	ID   int    "json:\"id\""
	Name string "json:\"name\""
}

func (*GroupCreate) createGroupRes() {}

// Ref: #/components/schemas/GroupList
type GroupList struct {
	ID   int    "json:\"id\""
	Name string "json:\"name\""
}

// Ref: #/components/schemas/GroupRead
type GroupRead struct {
	ID   int    "json:\"id\""
	Name string "json:\"name\""
}

func (*GroupRead) readGroupRes() {}

// Ref: #/components/schemas/GroupUpdate
type GroupUpdate struct {
	ID   int    "json:\"id\""
	Name string "json:\"name\""
}

func (*GroupUpdate) updateGroupRes() {}

// Ref: #/components/schemas/Group_UsersList
type GroupUsersList struct {
	ID   int    "json:\"id\""
	Age  int    "json:\"age\""
	Name string "json:\"name\""
}

type ListBattleOKApplicationJSON []BattleList

func (ListBattleOKApplicationJSON) listBattleRes() {}

type ListCarCarsOKApplicationJSON []CarCarsList

func (ListCarCarsOKApplicationJSON) listCarCarsRes() {}

type ListCarOKApplicationJSON []CarList

func (ListCarOKApplicationJSON) listCarRes() {}

type ListGroupOKApplicationJSON []GroupList

func (ListGroupOKApplicationJSON) listGroupRes() {}

type ListGroupUsersOKApplicationJSON []GroupUsersList

func (ListGroupUsersOKApplicationJSON) listGroupUsersRes() {}

type ListPetOKApplicationJSON []PetList

func (ListPetOKApplicationJSON) listPetRes() {}

type ListPokemonFightsOKApplicationJSON []PokemonFightsList

func (ListPokemonFightsOKApplicationJSON) listPokemonFightsRes() {}

type ListPokemonOKApplicationJSON []PokemonList

func (ListPokemonOKApplicationJSON) listPokemonRes() {}

type ListPokemonOpponentsOKApplicationJSON []PokemonOpponentsList

func (ListPokemonOpponentsOKApplicationJSON) listPokemonOpponentsRes() {}

type ListUserCarsOKApplicationJSON []UserCarsList

func (ListUserCarsOKApplicationJSON) listUserCarsRes() {}

type ListUserGroupsOKApplicationJSON []UserGroupsList

func (ListUserGroupsOKApplicationJSON) listUserGroupsRes() {}

type ListUserOKApplicationJSON []UserList

func (ListUserOKApplicationJSON) listUserRes() {}

// NewOptDateTime returns new OptDateTime with value set to v.
func NewOptDateTime(v time.Time) OptDateTime {
	return OptDateTime{
		Value: v,
		Set:   true,
	}
}

// OptDateTime is optional time.Time.
type OptDateTime struct {
	Value time.Time
	Set   bool
}

// IsSet returns true if OptDateTime was set.
func (o OptDateTime) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptDateTime) Reset() {
	var v time.Time
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptDateTime) SetTo(v time.Time) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptDateTime) Get() (v time.Time, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptDateTime) Or(d time.Time) time.Time {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptFloat64 returns new OptFloat64 with value set to v.
func NewOptFloat64(v float64) OptFloat64 {
	return OptFloat64{
		Value: v,
		Set:   true,
	}
}

// OptFloat64 is optional float64.
type OptFloat64 struct {
	Value float64
	Set   bool
}

// IsSet returns true if OptFloat64 was set.
func (o OptFloat64) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptFloat64) Reset() {
	var v float64
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptFloat64) SetTo(v float64) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptFloat64) Get() (v float64, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptFloat64) Or(d float64) float64 {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptInt returns new OptInt with value set to v.
func NewOptInt(v int) OptInt {
	return OptInt{
		Value: v,
		Set:   true,
	}
}

// OptInt is optional int.
type OptInt struct {
	Value int
	Set   bool
}

// IsSet returns true if OptInt was set.
func (o OptInt) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptInt) Reset() {
	var v int
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptInt) SetTo(v int) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptInt) Get() (v int, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptInt) Or(d int) int {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptString returns new OptString with value set to v.
func NewOptString(v string) OptString {
	return OptString{
		Value: v,
		Set:   true,
	}
}

// OptString is optional string.
type OptString struct {
	Value string
	Set   bool
}

// IsSet returns true if OptString was set.
func (o OptString) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptString) Reset() {
	var v string
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptString) SetTo(v string) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptString) Get() (v string, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// Ref: #/components/schemas/PetCreate
type PetCreate struct {
	ID   int    "json:\"id\""
	Name string "json:\"name\""
	Age  int    "json:\"age\""
}

func (*PetCreate) createPetRes() {}

// Ref: #/components/schemas/PetList
type PetList struct {
	ID   int    "json:\"id\""
	Name string "json:\"name\""
	Age  int    "json:\"age\""
}

// Ref: #/components/schemas/PetRead
type PetRead struct {
	ID   int    "json:\"id\""
	Name string "json:\"name\""
	Age  int    "json:\"age\""
}

func (*PetRead) readPetRes() {}

// Ref: #/components/schemas/PetUpdate
type PetUpdate struct {
	ID   int    "json:\"id\""
	Name string "json:\"name\""
	Age  int    "json:\"age\""
}

func (*PetUpdate) updatePetRes() {}

// Ref: #/components/schemas/PokemonCreate
type PokemonCreate struct {
	ID          int       "json:\"id\""
	Name        string    "json:\"name\""
	Description string    "json:\"description\""
	Weight      float64   "json:\"weight\""
	Height      float64   "json:\"height\""
	CreatedAt   time.Time "json:\"created_at\""
	UpdatedAt   time.Time "json:\"updated_at\""
}

func (*PokemonCreate) createPokemonRes() {}

// Ref: #/components/schemas/Pokemon_FightsList
type PokemonFightsList struct {
	ID        int       "json:\"id\""
	Result    string    "json:\"result\""
	CreatedAt time.Time "json:\"created_at\""
	UpdatedAt time.Time "json:\"updated_at\""
}

// Ref: #/components/schemas/PokemonList
type PokemonList struct {
	ID          int       "json:\"id\""
	Name        string    "json:\"name\""
	Description string    "json:\"description\""
	Weight      float64   "json:\"weight\""
	Height      float64   "json:\"height\""
	CreatedAt   time.Time "json:\"created_at\""
	UpdatedAt   time.Time "json:\"updated_at\""
}

// Ref: #/components/schemas/Pokemon_OpponentsList
type PokemonOpponentsList struct {
	ID        int       "json:\"id\""
	Result    string    "json:\"result\""
	CreatedAt time.Time "json:\"created_at\""
	UpdatedAt time.Time "json:\"updated_at\""
}

// Ref: #/components/schemas/PokemonRead
type PokemonRead struct {
	ID          int       "json:\"id\""
	Name        string    "json:\"name\""
	Description string    "json:\"description\""
	Weight      float64   "json:\"weight\""
	Height      float64   "json:\"height\""
	CreatedAt   time.Time "json:\"created_at\""
	UpdatedAt   time.Time "json:\"updated_at\""
}

func (*PokemonRead) readPokemonRes() {}

// Ref: #/components/schemas/PokemonUpdate
type PokemonUpdate struct {
	ID          int       "json:\"id\""
	Name        string    "json:\"name\""
	Description string    "json:\"description\""
	Weight      float64   "json:\"weight\""
	Height      float64   "json:\"height\""
	CreatedAt   time.Time "json:\"created_at\""
	UpdatedAt   time.Time "json:\"updated_at\""
}

func (*PokemonUpdate) updatePokemonRes() {}

type R400 struct {
	Code   int    "json:\"code\""
	Status string "json:\"status\""
	Errors jx.Raw "json:\"errors\""
}

func (*R400) createBattleRes()         {}
func (*R400) createCarRes()            {}
func (*R400) createGroupRes()          {}
func (*R400) createPetRes()            {}
func (*R400) createPokemonRes()        {}
func (*R400) createUserRes()           {}
func (*R400) deleteBattleRes()         {}
func (*R400) deleteCarRes()            {}
func (*R400) deleteGroupRes()          {}
func (*R400) deletePetRes()            {}
func (*R400) deletePokemonRes()        {}
func (*R400) deleteUserRes()           {}
func (*R400) listBattleRes()           {}
func (*R400) listCarCarsRes()          {}
func (*R400) listCarRes()              {}
func (*R400) listGroupRes()            {}
func (*R400) listGroupUsersRes()       {}
func (*R400) listPetRes()              {}
func (*R400) listPokemonFightsRes()    {}
func (*R400) listPokemonOpponentsRes() {}
func (*R400) listPokemonRes()          {}
func (*R400) listUserCarsRes()         {}
func (*R400) listUserGroupsRes()       {}
func (*R400) listUserRes()             {}
func (*R400) readBattleContenderRes()  {}
func (*R400) readBattleOponentRes()    {}
func (*R400) readBattleRes()           {}
func (*R400) readCarRes()              {}
func (*R400) readGroupRes()            {}
func (*R400) readPetRes()              {}
func (*R400) readPokemonRes()          {}
func (*R400) readUserRes()             {}
func (*R400) updateBattleRes()         {}
func (*R400) updateCarRes()            {}
func (*R400) updateGroupRes()          {}
func (*R400) updatePetRes()            {}
func (*R400) updatePokemonRes()        {}
func (*R400) updateUserRes()           {}

type R404 struct {
	Code   int    "json:\"code\""
	Status string "json:\"status\""
	Errors jx.Raw "json:\"errors\""
}

func (*R404) deleteBattleRes()         {}
func (*R404) deleteCarRes()            {}
func (*R404) deleteGroupRes()          {}
func (*R404) deletePetRes()            {}
func (*R404) deletePokemonRes()        {}
func (*R404) deleteUserRes()           {}
func (*R404) listBattleRes()           {}
func (*R404) listCarCarsRes()          {}
func (*R404) listCarRes()              {}
func (*R404) listGroupRes()            {}
func (*R404) listGroupUsersRes()       {}
func (*R404) listPetRes()              {}
func (*R404) listPokemonFightsRes()    {}
func (*R404) listPokemonOpponentsRes() {}
func (*R404) listPokemonRes()          {}
func (*R404) listUserCarsRes()         {}
func (*R404) listUserGroupsRes()       {}
func (*R404) listUserRes()             {}
func (*R404) readBattleContenderRes()  {}
func (*R404) readBattleOponentRes()    {}
func (*R404) readBattleRes()           {}
func (*R404) readCarRes()              {}
func (*R404) readGroupRes()            {}
func (*R404) readPetRes()              {}
func (*R404) readPokemonRes()          {}
func (*R404) readUserRes()             {}
func (*R404) updateBattleRes()         {}
func (*R404) updateCarRes()            {}
func (*R404) updateGroupRes()          {}
func (*R404) updatePetRes()            {}
func (*R404) updatePokemonRes()        {}
func (*R404) updateUserRes()           {}

type R409 struct {
	Code   int    "json:\"code\""
	Status string "json:\"status\""
	Errors jx.Raw "json:\"errors\""
}

func (*R409) createBattleRes()         {}
func (*R409) createCarRes()            {}
func (*R409) createGroupRes()          {}
func (*R409) createPetRes()            {}
func (*R409) createPokemonRes()        {}
func (*R409) createUserRes()           {}
func (*R409) deleteBattleRes()         {}
func (*R409) deleteCarRes()            {}
func (*R409) deleteGroupRes()          {}
func (*R409) deletePetRes()            {}
func (*R409) deletePokemonRes()        {}
func (*R409) deleteUserRes()           {}
func (*R409) listBattleRes()           {}
func (*R409) listCarCarsRes()          {}
func (*R409) listCarRes()              {}
func (*R409) listGroupRes()            {}
func (*R409) listGroupUsersRes()       {}
func (*R409) listPetRes()              {}
func (*R409) listPokemonFightsRes()    {}
func (*R409) listPokemonOpponentsRes() {}
func (*R409) listPokemonRes()          {}
func (*R409) listUserCarsRes()         {}
func (*R409) listUserGroupsRes()       {}
func (*R409) listUserRes()             {}
func (*R409) readBattleContenderRes()  {}
func (*R409) readBattleOponentRes()    {}
func (*R409) readBattleRes()           {}
func (*R409) readCarRes()              {}
func (*R409) readGroupRes()            {}
func (*R409) readPetRes()              {}
func (*R409) readPokemonRes()          {}
func (*R409) readUserRes()             {}
func (*R409) updateBattleRes()         {}
func (*R409) updateCarRes()            {}
func (*R409) updateGroupRes()          {}
func (*R409) updatePetRes()            {}
func (*R409) updatePokemonRes()        {}
func (*R409) updateUserRes()           {}

type R500 struct {
	Code   int    "json:\"code\""
	Status string "json:\"status\""
	Errors jx.Raw "json:\"errors\""
}

func (*R500) createBattleRes()         {}
func (*R500) createCarRes()            {}
func (*R500) createGroupRes()          {}
func (*R500) createPetRes()            {}
func (*R500) createPokemonRes()        {}
func (*R500) createUserRes()           {}
func (*R500) deleteBattleRes()         {}
func (*R500) deleteCarRes()            {}
func (*R500) deleteGroupRes()          {}
func (*R500) deletePetRes()            {}
func (*R500) deletePokemonRes()        {}
func (*R500) deleteUserRes()           {}
func (*R500) listBattleRes()           {}
func (*R500) listCarCarsRes()          {}
func (*R500) listCarRes()              {}
func (*R500) listGroupRes()            {}
func (*R500) listGroupUsersRes()       {}
func (*R500) listPetRes()              {}
func (*R500) listPokemonFightsRes()    {}
func (*R500) listPokemonOpponentsRes() {}
func (*R500) listPokemonRes()          {}
func (*R500) listUserCarsRes()         {}
func (*R500) listUserGroupsRes()       {}
func (*R500) listUserRes()             {}
func (*R500) readBattleContenderRes()  {}
func (*R500) readBattleOponentRes()    {}
func (*R500) readBattleRes()           {}
func (*R500) readCarRes()              {}
func (*R500) readGroupRes()            {}
func (*R500) readPetRes()              {}
func (*R500) readPokemonRes()          {}
func (*R500) readUserRes()             {}
func (*R500) updateBattleRes()         {}
func (*R500) updateCarRes()            {}
func (*R500) updateGroupRes()          {}
func (*R500) updatePetRes()            {}
func (*R500) updatePokemonRes()        {}
func (*R500) updateUserRes()           {}

type UpdateBattleReq struct {
	Result    OptString   "json:\"result\""
	UpdatedAt OptDateTime "json:\"updated_at\""
	Contender OptInt      "json:\"contender\""
	Oponent   OptInt      "json:\"oponent\""
}

type UpdateCarReq struct {
	Model        OptString   "json:\"model\""
	RegisteredAt OptDateTime "json:\"registered_at\""
	Cars         []int       "json:\"cars\""
}

type UpdateGroupReq struct {
	Name  OptString "json:\"name\""
	Users []int     "json:\"users\""
}

type UpdatePetReq struct {
	Name OptString "json:\"name\""
	Age  OptInt    "json:\"age\""
}

type UpdatePokemonReq struct {
	Name        OptString   "json:\"name\""
	Description OptString   "json:\"description\""
	Weight      OptFloat64  "json:\"weight\""
	Height      OptFloat64  "json:\"height\""
	UpdatedAt   OptDateTime "json:\"updated_at\""
	Fights      []int       "json:\"fights\""
	Opponents   []int       "json:\"opponents\""
}

type UpdateUserReq struct {
	Age    OptInt    "json:\"age\""
	Name   OptString "json:\"name\""
	Cars   []int     "json:\"cars\""
	Groups []int     "json:\"groups\""
}

// Ref: #/components/schemas/User_CarsList
type UserCarsList struct {
	ID           int       "json:\"id\""
	Model        string    "json:\"model\""
	RegisteredAt time.Time "json:\"registered_at\""
}

// Ref: #/components/schemas/UserCreate
type UserCreate struct {
	ID   int    "json:\"id\""
	Age  int    "json:\"age\""
	Name string "json:\"name\""
}

func (*UserCreate) createUserRes() {}

// Ref: #/components/schemas/User_GroupsList
type UserGroupsList struct {
	ID   int    "json:\"id\""
	Name string "json:\"name\""
}

// Ref: #/components/schemas/UserList
type UserList struct {
	ID   int    "json:\"id\""
	Age  int    "json:\"age\""
	Name string "json:\"name\""
}

// Ref: #/components/schemas/UserRead
type UserRead struct {
	ID   int    "json:\"id\""
	Age  int    "json:\"age\""
	Name string "json:\"name\""
}

func (*UserRead) readUserRes() {}

// Ref: #/components/schemas/UserUpdate
type UserUpdate struct {
	ID   int    "json:\"id\""
	Age  int    "json:\"age\""
	Name string "json:\"name\""
}

func (*UserUpdate) updateUserRes() {}
