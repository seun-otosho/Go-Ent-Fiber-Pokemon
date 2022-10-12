// Code generated by ogen, DO NOT EDIT.

package ogent

import (
	"context"

	"github.com/go-faster/jx"

	ht "github.com/ogen-go/ogen/http"
)

var _ Handler = UnimplementedHandler{}

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

// CreateBattle implements createBattle operation.
//
// Creates a new Battle and persists it to storage.
//
// POST /battles
func (UnimplementedHandler) CreateBattle(ctx context.Context, req CreateBattleReq) (r CreateBattleRes, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateCar implements createCar operation.
//
// Creates a new Car and persists it to storage.
//
// POST /cars
func (UnimplementedHandler) CreateCar(ctx context.Context, req CreateCarReq) (r CreateCarRes, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateGroup implements createGroup operation.
//
// Creates a new Group and persists it to storage.
//
// POST /groups
func (UnimplementedHandler) CreateGroup(ctx context.Context, req CreateGroupReq) (r CreateGroupRes, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateNote implements createNote operation.
//
// Creates a new Note and persists it to storage.
//
// POST /notes
func (UnimplementedHandler) CreateNote(ctx context.Context, req CreateNoteReq) (r CreateNoteRes, _ error) {
	return r, ht.ErrNotImplemented
}

// CreatePet implements createPet operation.
//
// Creates a new Pet and persists it to storage.
//
// POST /pets
func (UnimplementedHandler) CreatePet(ctx context.Context, req CreatePetReq) (r CreatePetRes, _ error) {
	return r, ht.ErrNotImplemented
}

// CreatePokemon implements createPokemon operation.
//
// Creates a new Pokemon and persists it to storage.
//
// POST /pokemons
func (UnimplementedHandler) CreatePokemon(ctx context.Context, req CreatePokemonReq) (r CreatePokemonRes, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateTodo implements createTodo operation.
//
// Creates a new Todo and persists it to storage.
//
// POST /todos
func (UnimplementedHandler) CreateTodo(ctx context.Context, req jx.Raw) (r CreateTodoRes, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateUser implements createUser operation.
//
// Creates a new User and persists it to storage.
//
// POST /users
func (UnimplementedHandler) CreateUser(ctx context.Context, req CreateUserReq) (r CreateUserRes, _ error) {
	return r, ht.ErrNotImplemented
}

// DBHealth implements DBHealth operation.
//
// Ping the database and report.
//
// GET /db-health
func (UnimplementedHandler) DBHealth(ctx context.Context) (r DBHealthRes, _ error) {
	return r, ht.ErrNotImplemented
}

// DeleteBattle implements deleteBattle operation.
//
// Deletes the Battle with the requested ID.
//
// DELETE /battles/{id}
func (UnimplementedHandler) DeleteBattle(ctx context.Context, params DeleteBattleParams) (r DeleteBattleRes, _ error) {
	return r, ht.ErrNotImplemented
}

// DeleteCar implements deleteCar operation.
//
// Deletes the Car with the requested ID.
//
// DELETE /cars/{id}
func (UnimplementedHandler) DeleteCar(ctx context.Context, params DeleteCarParams) (r DeleteCarRes, _ error) {
	return r, ht.ErrNotImplemented
}

// DeleteGroup implements deleteGroup operation.
//
// Deletes the Group with the requested ID.
//
// DELETE /groups/{id}
func (UnimplementedHandler) DeleteGroup(ctx context.Context, params DeleteGroupParams) (r DeleteGroupRes, _ error) {
	return r, ht.ErrNotImplemented
}

// DeleteNote implements deleteNote operation.
//
// Deletes the Note with the requested ID.
//
// DELETE /notes/{id}
func (UnimplementedHandler) DeleteNote(ctx context.Context, params DeleteNoteParams) (r DeleteNoteRes, _ error) {
	return r, ht.ErrNotImplemented
}

// DeletePet implements deletePet operation.
//
// Deletes the Pet with the requested ID.
//
// DELETE /pets/{id}
func (UnimplementedHandler) DeletePet(ctx context.Context, params DeletePetParams) (r DeletePetRes, _ error) {
	return r, ht.ErrNotImplemented
}

// DeletePokemon implements deletePokemon operation.
//
// Deletes the Pokemon with the requested ID.
//
// DELETE /pokemons/{id}
func (UnimplementedHandler) DeletePokemon(ctx context.Context, params DeletePokemonParams) (r DeletePokemonRes, _ error) {
	return r, ht.ErrNotImplemented
}

// DeleteTodo implements deleteTodo operation.
//
// Deletes the Todo with the requested ID.
//
// DELETE /todos/{id}
func (UnimplementedHandler) DeleteTodo(ctx context.Context, params DeleteTodoParams) (r DeleteTodoRes, _ error) {
	return r, ht.ErrNotImplemented
}

// DeleteUser implements deleteUser operation.
//
// Deletes the User with the requested ID.
//
// DELETE /users/{id}
func (UnimplementedHandler) DeleteUser(ctx context.Context, params DeleteUserParams) (r DeleteUserRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListBattle implements listBattle operation.
//
// List Battles.
//
// GET /battles
func (UnimplementedHandler) ListBattle(ctx context.Context, params ListBattleParams) (r ListBattleRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListCar implements listCar operation.
//
// List Cars.
//
// GET /cars
func (UnimplementedHandler) ListCar(ctx context.Context, params ListCarParams) (r ListCarRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListGroup implements listGroup operation.
//
// List Groups.
//
// GET /groups
func (UnimplementedHandler) ListGroup(ctx context.Context, params ListGroupParams) (r ListGroupRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListGroupUsers implements listGroupUsers operation.
//
// List attached Users.
//
// GET /groups/{id}/users
func (UnimplementedHandler) ListGroupUsers(ctx context.Context, params ListGroupUsersParams) (r ListGroupUsersRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListNote implements listNote operation.
//
// List Notes.
//
// GET /notes
func (UnimplementedHandler) ListNote(ctx context.Context, params ListNoteParams) (r ListNoteRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListPet implements listPet operation.
//
// List Pets.
//
// GET /pets
func (UnimplementedHandler) ListPet(ctx context.Context, params ListPetParams) (r ListPetRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListPokemon implements listPokemon operation.
//
// List Pokemons.
//
// GET /pokemons
func (UnimplementedHandler) ListPokemon(ctx context.Context, params ListPokemonParams) (r ListPokemonRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListPokemonFights implements listPokemonFights operation.
//
// List attached Fights.
//
// GET /pokemons/{id}/fights
func (UnimplementedHandler) ListPokemonFights(ctx context.Context, params ListPokemonFightsParams) (r ListPokemonFightsRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListPokemonOpponents implements listPokemonOpponents operation.
//
// List attached Opponents.
//
// GET /pokemons/{id}/opponents
func (UnimplementedHandler) ListPokemonOpponents(ctx context.Context, params ListPokemonOpponentsParams) (r ListPokemonOpponentsRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListTodo implements listTodo operation.
//
// List Todos.
//
// GET /todos
func (UnimplementedHandler) ListTodo(ctx context.Context, params ListTodoParams) (r ListTodoRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListUser implements listUser operation.
//
// List Users.
//
// GET /users
func (UnimplementedHandler) ListUser(ctx context.Context, params ListUserParams) (r ListUserRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListUserCars implements listUserCars operation.
//
// List attached Cars.
//
// GET /users/{id}/cars
func (UnimplementedHandler) ListUserCars(ctx context.Context, params ListUserCarsParams) (r ListUserCarsRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListUserGroups implements listUserGroups operation.
//
// List attached Groups.
//
// GET /users/{id}/groups
func (UnimplementedHandler) ListUserGroups(ctx context.Context, params ListUserGroupsParams) (r ListUserGroupsRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadBattle implements readBattle operation.
//
// Finds the Battle with the requested ID and returns it.
//
// GET /battles/{id}
func (UnimplementedHandler) ReadBattle(ctx context.Context, params ReadBattleParams) (r ReadBattleRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadBattleContender implements readBattleContender operation.
//
// Find the attached Pokemon of the Battle with the given ID.
//
// GET /battles/{id}/contender
func (UnimplementedHandler) ReadBattleContender(ctx context.Context, params ReadBattleContenderParams) (r ReadBattleContenderRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadBattleOponent implements readBattleOponent operation.
//
// Find the attached Pokemon of the Battle with the given ID.
//
// GET /battles/{id}/oponent
func (UnimplementedHandler) ReadBattleOponent(ctx context.Context, params ReadBattleOponentParams) (r ReadBattleOponentRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadCar implements readCar operation.
//
// Finds the Car with the requested ID and returns it.
//
// GET /cars/{id}
func (UnimplementedHandler) ReadCar(ctx context.Context, params ReadCarParams) (r ReadCarRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadGroup implements readGroup operation.
//
// Finds the Group with the requested ID and returns it.
//
// GET /groups/{id}
func (UnimplementedHandler) ReadGroup(ctx context.Context, params ReadGroupParams) (r ReadGroupRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadNote implements readNote operation.
//
// Finds the Note with the requested ID and returns it.
//
// GET /notes/{id}
func (UnimplementedHandler) ReadNote(ctx context.Context, params ReadNoteParams) (r ReadNoteRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadPet implements readPet operation.
//
// Finds the Pet with the requested ID and returns it.
//
// GET /pets/{id}
func (UnimplementedHandler) ReadPet(ctx context.Context, params ReadPetParams) (r ReadPetRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadPokemon implements readPokemon operation.
//
// Finds the Pokemon with the requested ID and returns it.
//
// GET /pokemons/{id}
func (UnimplementedHandler) ReadPokemon(ctx context.Context, params ReadPokemonParams) (r ReadPokemonRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadTodo implements readTodo operation.
//
// Finds the Todo with the requested ID and returns it.
//
// GET /todos/{id}
func (UnimplementedHandler) ReadTodo(ctx context.Context, params ReadTodoParams) (r ReadTodoRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadUser implements readUser operation.
//
// Finds the User with the requested ID and returns it.
//
// GET /users/{id}
func (UnimplementedHandler) ReadUser(ctx context.Context, params ReadUserParams) (r ReadUserRes, _ error) {
	return r, ht.ErrNotImplemented
}

// UpdateBattle implements updateBattle operation.
//
// Updates a Battle and persists changes to storage.
//
// PATCH /battles/{id}
func (UnimplementedHandler) UpdateBattle(ctx context.Context, req UpdateBattleReq, params UpdateBattleParams) (r UpdateBattleRes, _ error) {
	return r, ht.ErrNotImplemented
}

// UpdateCar implements updateCar operation.
//
// Updates a Car and persists changes to storage.
//
// PATCH /cars/{id}
func (UnimplementedHandler) UpdateCar(ctx context.Context, req UpdateCarReq, params UpdateCarParams) (r UpdateCarRes, _ error) {
	return r, ht.ErrNotImplemented
}

// UpdateGroup implements updateGroup operation.
//
// Updates a Group and persists changes to storage.
//
// PATCH /groups/{id}
func (UnimplementedHandler) UpdateGroup(ctx context.Context, req UpdateGroupReq, params UpdateGroupParams) (r UpdateGroupRes, _ error) {
	return r, ht.ErrNotImplemented
}

// UpdateNote implements updateNote operation.
//
// Updates a Note and persists changes to storage.
//
// PATCH /notes/{id}
func (UnimplementedHandler) UpdateNote(ctx context.Context, req UpdateNoteReq, params UpdateNoteParams) (r UpdateNoteRes, _ error) {
	return r, ht.ErrNotImplemented
}

// UpdatePet implements updatePet operation.
//
// Updates a Pet and persists changes to storage.
//
// PATCH /pets/{id}
func (UnimplementedHandler) UpdatePet(ctx context.Context, req UpdatePetReq, params UpdatePetParams) (r UpdatePetRes, _ error) {
	return r, ht.ErrNotImplemented
}

// UpdatePokemon implements updatePokemon operation.
//
// Updates a Pokemon and persists changes to storage.
//
// PATCH /pokemons/{id}
func (UnimplementedHandler) UpdatePokemon(ctx context.Context, req UpdatePokemonReq, params UpdatePokemonParams) (r UpdatePokemonRes, _ error) {
	return r, ht.ErrNotImplemented
}

// UpdateTodo implements updateTodo operation.
//
// Updates a Todo and persists changes to storage.
//
// PATCH /todos/{id}
func (UnimplementedHandler) UpdateTodo(ctx context.Context, req jx.Raw, params UpdateTodoParams) (r UpdateTodoRes, _ error) {
	return r, ht.ErrNotImplemented
}

// UpdateUser implements updateUser operation.
//
// Updates a User and persists changes to storage.
//
// PATCH /users/{id}
func (UnimplementedHandler) UpdateUser(ctx context.Context, req UpdateUserReq, params UpdateUserParams) (r UpdateUserRes, _ error) {
	return r, ht.ErrNotImplemented
}
