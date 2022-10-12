package handlers

//app.Get("/all", func(c *fiber.Ctx) error {
//	pokemons, err := client.Pokemon.Query().All(ctx)
//	if err != nil {
//		return fmt.Errorf("failed querying pokemons: %w", err)
//	}
//	log.Println("returned pokemons:", pokemons)
//
//	return c.Status(200).JSON(pokemons)
//})
//
//app.Post("/create", func(c *fiber.Ctx) error {
//	payload := struct {
//		Name        string  `json:"name"`
//		Description string  `json:"description"`
//		Weight      float64 `json:"weight"`
//		Height      float64 `json:"height"`
//	}{}
//
//	if err := c.BodyParser(&payload); err != nil {
//		return err
//	}
//
//	pokemon, err := client.Pokemon.
//		Create().
//		SetName(payload.Name).
//		SetDescription(payload.Description).
//		SetWeight(payload.Weight).
//		SetHeight(payload.Height).
//		Save(ctx)
//	if err != nil {
//		return fmt.Errorf("failed creating pokemon: %w", err)
//	}
//	log.Println("pokemon created: ", pokemon)
//	return c.Status(200).JSON(pokemon)
//})
//
//app.Delete("/delete", func(c *fiber.Ctx) error {
//
//	payload := struct {
//		Id int `json:"id"`
//	}{}
//
//	if err := c.BodyParser(&payload); err != nil {
//		return err
//	}
//
//	err := client.Pokemon.DeleteOneID(payload.Id).Exec(ctx)
//	if err != nil {
//		return fmt.Errorf("failed deleting pokemon: %w", err)
//	}
//	return c.Status(200).SendString("Success")
//})
//
//app.Put("/update", func(c *fiber.Ctx) error {
//	payload := struct {
//		Id          int     `json:"id"`
//		Name        string  `json:"name"`
//		Description string  `json:"description"`
//		Weight      float64 `json:"weight"`
//		Height      float64 `json:"height"`
//	}{}
//
//	if err := c.BodyParser(&payload); err != nil {
//		return err
//	}
//
//	pokemon, err := client.Pokemon.
//		UpdateOneID(payload.Id).
//		SetName(payload.Name).
//		SetDescription(payload.Description).
//		SetWeight(payload.Weight).
//		SetHeight(payload.Height).
//		Save(ctx)
//	if err != nil {
//		return fmt.Errorf("failed updating pokemon: %w", err)
//	}
//	return c.Status(200).JSON(pokemon)
//})
//
//app.Post("/create/battle", func(c *fiber.Ctx) error {
//	payload := struct {
//		Result    string `json:"result"`
//		Contender int    `json:"contender"`
//		Oponent   int    `json:"oponent"`
//	}{}
//
//	if err := c.BodyParser(&payload); err != nil {
//		return err
//	}
//
//	battle, err := client.Battle.
//		Create().
//		SetResult(payload.Result).
//		SetContenderID(payload.Contender).
//		SetOponentID(payload.Oponent).
//		Save(ctx)
//	if err != nil {
//		return fmt.Errorf("failed creating battle: %w", err)
//	}
//	log.Println("battle created: ", battle)
//	return c.Status(200).JSON(battle)
//})
//
//app.Get("/all/battle", func(c *fiber.Ctx) error {
//	battles, err := client.Battle.Query().WithContender().WithOponent().All(ctx)
//	if err != nil {
//		return fmt.Errorf("failed querying battles: %w", err)
//	}
//	log.Println("returned battles:", battles)
//
//	return c.Status(200).JSON(battles)
//})
