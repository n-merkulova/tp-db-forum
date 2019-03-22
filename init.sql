CREATE EXTENSION IF NOT EXISTS citext;

CREATE TABLE IF NOT EXISTS Users
(
    nickname CITEXT NOT NULL PRIMARY KEY,
    fullname TEXT NOT NULL,
    email CITEXT UNIQUE NOT NULL,
    about TEXT
);

-- CREATE UNIQUE INDEX IF NOT EXISTS index_users_email ON Users(email);

CREATE TABLE IF NOT EXISTS Forum
(
    slug CITEXT NOT NULL PRIMARY KEY,
    forumUser CITEXT REFERENCES Users(nickname) NOT NULL,
    title CITEXT NOT NULL,
    posts INTEGER DEFAULT 0,
    threads INTEGER DEFAULT 0
);

CREATE TABLE IF NOT EXISTS Thread
(
    id SERIAL NOT NULL PRIMARY KEY,
    author CITEXT REFERENCES Users(nickname) NOT NULL,
    forum CITEXT REFERENCES Forum(slug),
    votes INTEGER DEFAULT 0,
    slug CITEXT UNIQUE,
    title TEXT NOT NULL,
    message TEXT NOT NULL,
    created TIMESTAMP WITH TIME ZONE DEFAULT now()
);

CREATE TABLE IF NOT EXISTS Post
(
    id SERIAL PRIMARY KEY ,
    parent INTEGER DEFAULT 0,
    author CITEXT NOT NULL,
    message TEXT NOT NULL,
    isEdited BOOLEAN,
    forum CITEXT,
    created TIMESTAMP WITH TIME ZONE DEFAULT now(),
    thread INTEGER,

    path INTEGER[] NOT NULL
);

CREATE TABLE IF NOT EXISTS Vote (
    id SERIAL NOT NULL PRIMARY KEY,
    nickname CITEXT REFERENCES Users(nickname) NOT NULL,
    threadID INTEGER NOT NULL,
    voice INT NOT NULL,
);
