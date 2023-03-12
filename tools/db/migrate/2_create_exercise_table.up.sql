CREATE TABLE exercise_categories
(
    id          INT(11) PRIMARY KEY AUTO_INCREMENT COMMENT 'カテゴリーID',
    user_id     INT(11)     NOT NULL COMMENT 'ユーザーID',
    name        VARCHAR(255) NOT NULL COMMENT 'カテゴリー名' CHECK (LENGTH(`name`) > 0),
    description TEXT        NULL COMMENT 'カテゴリーの説明',
    created_at  TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'レコード作成日時',
    updated_at  TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'レコード更新日時',
    deleted_at  TIMESTAMP   NULL     DEFAULT NULL COMMENT 'レコード削除日時',
    FOREIGN KEY (user_id) REFERENCES users (id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT ='種目カテゴリマスタ';

CREATE TABLE exercises
(
    id            INT(11) PRIMARY KEY AUTO_INCREMENT COMMENT '種目ID',
    name VARCHAR(255) NOT NULL COMMENT '種目名' CHECK (LENGTH(name) > 0),
    description   TEXT COMMENT '種目の説明',
    user_id       INT(11)     NOT NULL COMMENT 'ユーザーID',
    category_id   INT(11)     NOT NULL COMMENT 'カテゴリーID',
    created_at    TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'レコード作成日時',
    updated_at    TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'レコード更新日時',
    deleted_at    TIMESTAMP   NULL     DEFAULT NULL COMMENT 'レコード削除日時',
    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (category_id) REFERENCES exercise_categories (id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT ='種目マスタ';

CREATE TABLE training_records
(
    id         INT(11) PRIMARY KEY COMMENT '記録ID',
    user_id    INT(11)   NOT NULL COMMENT 'ユーザーID',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'レコード作成日時',
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'レコード更新日時',
    deleted_at TIMESTAMP NULL     DEFAULT NULL COMMENT 'レコード削除日時',
    FOREIGN KEY (user_id) REFERENCES users (id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT ='トレーニング記録';

CREATE TABLE training_sets
(
    id                 INT(11)   NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT 'セットID',
    training_record_id INT(11)   NOT NULL COMMENT '記録ID',
    exercise_id        INT(11)   NOT NULL COMMENT '種目ID',
    memo               TEXT COMMENT 'その他詳細',
    created_at         TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'レコード作成日時',
    updated_at         TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'レコード更新日時',
    deleted_at         TIMESTAMP NULL     DEFAULT NULL COMMENT 'レコード削除日時',
    FOREIGN KEY (training_record_id) REFERENCES training_records (id),
    FOREIGN KEY (exercise_id) REFERENCES exercises (id),
    UNIQUE KEY ix_training_training_record_sets_record_id_exercise_id (training_record_id, exercise_id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT ='トレーニングセット';

CREATE TABLE training_set_details
(
    id                 INT(11)   NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT 'トレーニング詳細ID',
    training_record_id INT(11)   NOT NULL COMMENT '記録ID',
    training_set_id    INT(11)   NOT NULL COMMENT 'トレーニングセットID',
    weight             INT(11)   NOT NULL COMMENT '重量(kg): NULLの場合自重',
    reps               INT(11)   NOT NULL COMMENT 'レップ数',
    memo               TEXT COMMENT 'その他詳細',
    created_at         TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'レコード作成日時',
    updated_at         TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'レコード更新日時',
    deleted_at         TIMESTAMP NULL     DEFAULT NULL COMMENT 'レコード削除日時',
    FOREIGN KEY (training_record_id) REFERENCES training_records (id),
    FOREIGN KEY (training_set_id) REFERENCES training_sets (id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT ='トレーニングセット詳細';
