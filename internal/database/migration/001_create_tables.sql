-- Write your migrate up statements here

ALTER DATABASE "vocab-booster" SET timezone TO 'UTC';

CREATE SCHEMA IF NOT EXISTS public;

-- VOCABULARY --

CREATE TABLE vocabularies (
    id TEXT PRIMARY KEY,
    author_id TEXT NOT NULL,
    term VARCHAR(30) NOT NULL,
    parts_of_speech TEXT[] NOT NULL,
    ipa VARCHAR(30) NOT NULL,
    audio VARCHAR(50) NOT NULL,
    synonyms TEXT[] NOT NULL DEFAULT '{}',
    antonyms TEXT[] NOT NULL DEFAULT '{}',
    created_at TIMESTAMPTZ DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT NOW() NOT NULL,

    CONSTRAINT vocabularies_unique_term
        UNIQUE (term)
);

CREATE TABLE vocabulary_examples (
    id TEXT PRIMARY KEY,
    vocabulary_id TEXT NOT NULL,
    audio VARCHAR(50) NOT NULL,
    content JSONB NOT NULL, -- Multilingual
    main_word JSONB NOT NULL,
    pos_tags JSONB NOT NULL DEFAULT '{}',
    sentiment JSONB NOT NULL,
    dependencies JSONB NOT NULL DEFAULT '{}',
    verbs JSONB NOT NULL DEFAULT '{}',
    level VARCHAR(20) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW() NOT NULL,

    FOREIGN KEY(vocabulary_id) REFERENCES vocabularies(id)
);

CREATE INDEX idx_vocabulary_examples_vocabulary_id_created_at ON vocabulary_examples(vocabulary_id, created_at DESC);

CREATE TABLE verb_conjugations (
    id TEXT PRIMARY KEY,
    vocabulary_id TEXT NOT NULL,
    value VARCHAR(30) NOT NULL,
    base VARCHAR(30) NOT NULL,
    form VARCHAR(30) NOT NULL,

    FOREIGN KEY(vocabulary_id) REFERENCES vocabularies(id),

    UNIQUE (value, form)
);

CREATE TABLE vocabulary_scraping_items (
    id TEXT PRIMARY KEY,
    term VARCHAR(30) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW() NOT NULL,

    UNIQUE (term)
);

-- COMMUNITY --

CREATE TABLE community_sentences (
    id TEXT PRIMARY KEY,
    user_id TEXT NOT NULL,
    vocabulary_id TEXT NOT NULL,
    content JSONB NOT NULL, -- Multilingual
    required_vocabulary TEXT[] NOT NULL DEFAULT '{}',
    required_tense TEXT NOT NULL,
    sentiment JSONB NOT NULL,
    clauses JSONB NOT NULL DEFAULT '{}',
    pos_tags JSONB NOT NULL DEFAULT '{}',
    dependencies JSONB NOT NULL DEFAULT '{}',
    verbs JSONB NOT NULL DEFAULT '{}',
    level VARCHAR(20) NOT NULL,
    stats_like int NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW() NOT NULL,

    FOREIGN KEY(vocabulary_id) REFERENCES vocabularies(id)
);

CREATE INDEX idx_comm_sent_user_id_created_at ON community_sentences(user_id, created_at DESC);
CREATE INDEX idx_comm_sent_vocabulary_id_created_at ON community_sentences(vocabulary_id, created_at DESC);

CREATE TABLE community_sentence_drafts (
    id TEXT PRIMARY KEY,
    user_id TEXT NOT NULL,
    vocabulary_id TEXT NOT NULL,
    content JSONB NOT NULL, -- Multilingual
    required_vocabulary TEXT[] NOT NULL DEFAULT '{}',
    required_tense TEXT NOT NULL,
    is_correct BOOLEAN NOT NULL,
    grammar_errors JSONB NOT NULL DEFAULT '{}',
    error_code VARCHAR(30) NOT NULL,
    sentiment JSONB NOT NULL,
    clauses JSONB NOT NULL DEFAULT '{}',
    created_at TIMESTAMPTZ DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT NOW() NOT NULL,

    FOREIGN KEY(vocabulary_id) REFERENCES vocabularies(id)
);

CREATE INDEX idx_comm_sent_dra_user_id_created_at ON community_sentence_drafts(user_id, created_at DESC);
CREATE INDEX idx_comm_sent_dra_vocabulary_id_created_at ON community_sentence_drafts(vocabulary_id, created_at DESC);
CREATE INDEX idx_comm_sent_dra_vocabulary_id_user_id ON community_sentence_drafts(vocabulary_id, user_id);

CREATE TABLE community_sentence_likes (
    id TEXT NOT NULL,
    user_id TEXT NOT NULL,
    sentence_id TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW() NOT NULL,

    PRIMARY KEY (user_id, sentence_id),
    FOREIGN KEY(sentence_id) REFERENCES community_sentences(id)
);

CREATE INDEX idx_community_sentence_likes_user_id_sentence_id ON community_sentence_likes(user_id, sentence_id);

-- EXERCISE --

CREATE TABLE exercises (
    id TEXT PRIMARY KEY,
    vocabulary_example_id TEXT NOT NULL,
    audio VARCHAR(50) NOT NULL,
    level VARCHAR(20) NOT NULL,
    content JSONB NOT NULL, -- Multilingual
    vocabulary VARCHAR(30) NOT NULL,
    correct_answer VARCHAR(30) NOT NULL,
    options TEXT[] NOT NULL DEFAULT '{}',
    created_at TIMESTAMPTZ DEFAULT NOW() NOT NULL,

    FOREIGN KEY(vocabulary_example_id) REFERENCES vocabulary_examples(id)
);

CREATE INDEX idx_exercises_vocabulary_example_id ON exercises(vocabulary_example_id);
CREATE INDEX idx_exercises_vocabulary_created_at ON exercises(vocabulary, created_at DESC);
CREATE INDEX idx_exercises_level_created_at ON exercises(level, created_at DESC);


CREATE TABLE user_exercise_statuses (
    id TEXT PRIMARY KEY,
    exercise_id TEXT NOT NULL,
    user_id TEXT NOT NULL,
    correct_streak INT NOT NULL,
    answer_count INT NOT NULL,
    is_favorite BOOLEAN NOT NULL,
    is_mastered BOOLEAN NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT NOW() NOT NULL,
    next_review_at TIMESTAMPTZ NOT NULL,

    FOREIGN KEY(exercise_id) REFERENCES exercises(id),

    UNIQUE (user_id, exercise_id)
);

CREATE INDEX idx_user_exercise_statuses_user_id_updated_at ON user_exercise_statuses(user_id, updated_at DESC);

CREATE TABLE user_exercise_interacted_histories (
    id TEXT PRIMARY KEY,
    exercise_id TEXT NOT NULL,
    user_id TEXT NOT NULL,
    date TIMESTAMPTZ NOT NULL,

    FOREIGN KEY(exercise_id) REFERENCES exercises(id),
    UNIQUE (user_id, exercise_id, date)
);

CREATE INDEX idx_user_exercise_interacted_histories_user_id_exercise_id_date ON user_exercise_interacted_histories(user_id, exercise_id, date);

-- EXERCISE COLLECTION --

CREATE TABLE exercise_collections (
    id TEXT PRIMARY KEY,
    name JSONB NOT NULL, -- Multilingual
    slug TEXT NOT NULL,
    criteria TEXT NOT NULL DEFAULT '',
    is_from_system BOOLEAN NOT NULL DEFAULT FALSE,
    "order" INTEGER NOT NULL DEFAULT 0,
    image TEXT NOT NULL DEFAULT '',
    stats_exercises INTEGER NOT NULL DEFAULT 0,
    last_stats_updated_at TIMESTAMPTZ DEFAULT NOW() NOT NULL,

    UNIQUE (slug)
);

CREATE INDEX idx_exercise_collections_slug ON exercise_collections(slug);

CREATE TABLE user_exercise_collection_status (
    id TEXT NOT NULL,
    user_id TEXT NOT NULL,
    collection_id  TEXT NOT NULL,
    interacted_exercises INTEGER NOT NULL DEFAULT 0,

    PRIMARY KEY (user_id, collection_id)
);

CREATE INDEX idx_user_exercise_collections_user_id_collection_id ON user_exercise_collection_status(user_id, collection_id);

---- create above / drop below ----

-- Write your migrate down statements here. If this migration is irreversible

