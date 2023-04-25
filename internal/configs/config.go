package configs

type Config struct {
	Server          Server          `yaml:"Server"`
	Postgres        Postgres        `yaml:"Postgres"`
	Cors            Cors            `yaml:"Cors"`
	Redis           Redis           `yaml:"Redis"`
	Minio           Minio           `yaml:"Minio"`
	Kafka           Kafka           `yaml:"Kafka"`
	ChatsService    ChatsService    `yaml:"ChatsService"`
	UsersService    UsersService    `yaml:"UsersService"`
	MessagesService MessagesService `yaml:"MessagesService"`
}

type Server struct {
	Port string `yaml:"port"`
}

type Postgres struct {
	DB             string `yaml:"db"`
	ConnectionToDB string `yaml:"connectionToDB"`
}

type Cors struct {
	AllowMethods     []string `yaml:"allowMethods"`
	AllowOrigins     []string `yaml:"allowOrigins"`
	AllowHeaders     []string `yaml:"allowHeaders"`
	ExposeHeaders    []string `yaml:"exposeHeaders"`
	AllowCredentials bool     `yaml:"allowCredentials"`
}

type Redis struct {
	Addr string `yaml:"addr"`
}

type Minio struct {
	Endpoint string `yaml:"endpoint"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Token    string `yaml:"token"`
}

type Kafka struct {
	BrokerList []string `yaml:"brokerList"`
	GroupID    string   `yaml:"groupID"`
}

type ChatsService struct {
	Addr string `yaml:"addr"`
}

type UsersService struct {
	Addr string `yaml:"addr"`
}

type MessagesService struct {
	Addr string `yaml:"addr"`
}

const DefaultAvatarUrl = `https://avatars.mds.yandex.net/i?id=fb89295056d345e663a7c3c998a0dfd44ea37174-8497272-images-thumbs&n=13&exp=1`

const Chat = 0
const Group = 1
const Channel = 2

const Create = 0
const Edit = 1
const Delete = 2
