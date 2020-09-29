package main

import (
	"github.com/Kamva/mgm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"go.mongodb.org/mongo-driver/mongo/options"
	"user/user"
)

func init() {
	_ = mgm.SetDefaultConfig(nil, "check_echo", options.Client().ApplyURI("mongodb+srv://sdw:1q2w3e!%40#@cluster0.cl17y.mongodb.net/test?retryWrites=true&w=majority"))
}

func main() {
	e := echo.New()
	e.Debug = true

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	sample_echo := e.Group("/user")
	{
		sample_echo.POST("", user.Create)
	}

	e.Logger.Fatal(e.Start(":1323"))
}

/*

func main() {
	clientOptions := options.Client().ApplyURI("mongodb+srv://sdw:1q2w3e!%40#@cluster0.cl17y.mongodb.net/test?retryWrites=true&w=majority")

	//fmt.Println("clientOptions TYPE : ", reflect.TypeOf(clientOptions), "\n")

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		//fmt.Println("mongo.Connect() ERROR : ", err)
		os.Exit(1)
	}

	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	col := client.Database("test").Collection("user")
	//fmt.Println("Collection type : ", reflect.TypeOf(col), "\n")

	oneDoc := MongoFields{
		FeildStr: "test2",
		FeildInt: 32,
	}

	add(col, ctx, oneDoc)

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})

	e.GET("/users/:id", getUser)

	e.GET("/show", show)

	e.GET("/setUser", setUser)

	e.Logger.Fatal(e.Start(":1323"))
}

type MongoFields struct {
	FeildStr string
	FeildInt int
}

type User struct {
	name string
	age  string
}

func add(col *mongo.Collection, ctx context.Context, oneDoc MongoFields) {
	result, insertErr := col.InsertOne(ctx, oneDoc)

	if insertErr != nil {
		fmt.Println("InsertOne  ERROR : ", insertErr)
		os.Exit(1)
	} else {
		fmt.Println("InsertOne() result type : ", reflect.TypeOf(result))
		newID := result.InsertedID
		fmt.Println("InsertOne() newID : ", newID)
	}
}

func getUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func setUser(c echo.Context) error {
	user := User{
		name: c.QueryParam("name"),
		age:  c.QueryParam("age"),
	}

	clientOptions := options.Client().ApplyURI("mongodb+srv://sdw:1q2w3e!%40#@cluster0.cl17y.mongodb.net/test?retryWrites=true&w=majority")

	//fmt.Println("clientOptions TYPE : ", reflect.TypeOf(clientOptions), "\n")

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		//fmt.Println("mongo.Connect() ERROR : ", err)
		os.Exit(1)
	}

	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	col := client.Database("test").Collection("user")

	return user
}

func show(c echo.Context) error {
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:"+team+", member:"+member)
}
*/
