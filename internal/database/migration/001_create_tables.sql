-- Write your migrate up statements here

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
    updated_at TIMESTAMPTZ DEFAULT NOW() NOT NULL

    CONSTRAINT vocabularies_unique_term
        UNIQUE (term)
);

CREATE TABLE vocabulary_examples (
    id TEXT PRIMARY KEY,
    vocabulary_id TEXT NOT NULL,
    audio VARCHAR(50) NOT NULL,
    content TEXT NOT NULL,
    translated JSONB NOT NULL,
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
    content TEXT NOT NULL,
    required_vocabulary TEXT[] NOT NULL DEFAULT '{}',
    required_tense TEXT NOT NULL,
    translated JSONB NOT NULL,
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
    content TEXT NOT NULL,
    required_vocabulary TEXT[] NOT NULL DEFAULT '{}',
    required_tense TEXT NOT NULL,
    is_correct BOOLEAN NOT NULL,
    is_grammar_correct BOOLEAN NOT NULL,
    grammar_errors JSONB NOT NULL DEFAULT '{}',
    is_english BOOLEAN NOT NULL,
    is_tense_correct BOOLEAN NOT NULL,
    is_vocabulary_correct BOOLEAN NOT NULL,
    translated JSONB NOT NULL,
    sentiment JSONB NOT NULL,
    clauses JSONB NOT NULL DEFAULT '{}',
    created_at TIMESTAMPTZ DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT NOW() NOT NULL,

    FOREIGN KEY(vocabulary_id) REFERENCES vocabularies(id)
);

CREATE INDEX idx_comm_sent_dra_user_id_created_at ON community_sentence_drafts(user_id, created_at DESC);
CREATE INDEX idx_comm_sent_dra_vocabulary_id_created_at ON community_sentence_drafts(vocabulary_id, created_at DESC);
CREATE INDEX idx_comm_sent_dra_vocabulary_id_user_id ON community_sentence_drafts(vocabulary_id, user_id);

-- USER VOCABULARY --

CREATE TABLE user_vocabularies (
    id TEXT PRIMARY KEY,
    user_id TEXT NOT NULL,
    vocabulary_id TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW() NOT NULL,

    PRIMARY KEY (user_id, vocabulary_id),
    FOREIGN KEY (vocabulary_id) REFERENCES vocabularies(id)
);

CREATE INDEX idx_user_vocabularies_user_id_created_at ON user_vocabularies(user_id, created_at DESC);

-- VOCABULARY COLLECTION --

CREATE TABLE collections (
    id TEXT PRIMARY KEY,
    user_id TEXT NOT NULL,
    name varchar(30) NOT NULL,
    description TEXT NOT NULL,
    num_of_vocabulary INT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT NOW() NOT NULL
);

CREATE INDEX idx_collections_user_id_updated_at ON collections(user_id, updated_at);

CREATE TABLE collection_and_vocabularies (
    collection_id TEXT NOT NULL,
    vocabulary_id TEXT NOT NULL,
    value VARCHAR(30) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW() NOT NULL,
    PRIMARY KEY (collection_id, vocabulary_id),
    FOREIGN KEY (collection_id) REFERENCES collections(id) ON DELETE CASCADE,
    FOREIGN KEY (vocabulary_id) REFERENCES vocabularies(id) ON DELETE CASCADE
);

CREATE INDEX idx_collection_and_vocabularies_collection_id_created_at ON collection_and_vocabularies(collection_id, vocabulary_id);

-- EXERCISE --

CREATE TABLE exercises (
    id TEXT PRIMARY KEY,
    vocabulary_example_id TEXT NOT NULL,
    audio VARCHAR(50) NOT NULL,
    level VARCHAR(20) NOT NULL,
    content TEXT NOT NULL,
    translated JSONB NOT NULL,
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

---- create above / drop below ----

-- Write your migrate down statements here. If this migration is irreversible
DROP TABLE exercises CASCADE;
DROP TABLE user_exercise_statuses CASCADE;
