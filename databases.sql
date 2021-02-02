CREATE TABLE IF NOT EXISTS users(
        id INTEGER PRIMARY KEY,
        username TEXT NOT NULL UNIQUE,
        password TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS scores(
        id INTEGER PRIMARY KEY,
        score TEXT NOT NULL,
        date TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS users_scores(
    user_id INTEGER NOT NULL, 
    score_id INTEGER NOT NULL UNIQUE,
    
    FOREIGN KEY (user_id)
        REFERENCES users (id)
            ON UPDATE CASCADE
            ON DELETE CASCADE,

    FOREIGN KEY(score_id)
        REFERENCES scores (id)
            ON UPDATE CASCADE
            ON DELETE CASCADE
);

INSERT INTO users (username,password) 
    VALUES
        ('cesar','cfabrica46'),
        ('arturo','01234');

INSERT INTO scores (score,date) 
    VALUES
        (20,datetime('now')),
        (15,datetime('now')),
        (25,datetime('now'));

INSERT INTO users_scores (user_id,score_id)
    VALUES
        (1,1),
        (1,2),
        (2,3);