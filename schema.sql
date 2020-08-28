CREATE SEQUENCE writer_role_id_seq;
CREATE SEQUENCE writers_id_seq;
CREATE SEQUENCE publishers_id_seq;
CREATE SEQUENCE books_id_seq;
CREATE SEQUENCE book_writer_association_id_seq;

CREATE TABLE writer_role (
  id INT NOT NULL DEFAULT nextval('writers_id_seq'::regclass),
  name VARCHAR(50),
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  version_number INT NOT NULL DEFAULT 1
);

COMMENT ON TABLE writer_role IS '著者役割';
COMMENT ON COLUMN writer_role.id IS 'ID';
COMMENT ON COLUMN writer_role.name IS '著者役割名';
COMMENT ON COLUMN writer_role.created_at IS '作成日時';
COMMENT ON COLUMN writer_role.updated_at IS '更新日時';
COMMENT ON COLUMN writer_role.version_number IS 'バージョン番号';

CREATE TABLE writers (
  id INT NOT NULL DEFAULT nextval('writers_id_seq'::regclass),
  family_name VARCHAR(50),
  given_name VARCHAR(50),
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  version_number INT NOT NULL DEFAULT 1
);

COMMENT ON TABLE writers IS '著者';
COMMENT ON COLUMN writers.id IS 'ID';
COMMENT ON COLUMN writers.family_name IS '姓';
COMMENT ON COLUMN writers.given_name IS '名';
COMMENT ON COLUMN writers.created_at IS '作成日時';
COMMENT ON COLUMN writers.updated_at IS '更新日時';
COMMENT ON COLUMN writers.version_number IS 'バージョン番号';

CREATE TABLE publishers (
  id INT NOT NULL DEFAULT nextval('publishers_id_seq'::regclass),
  name VARCHAR(100),
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  version_number INT NOT NULL DEFAULT 1
);

COMMENT ON TABLE publishers IS '出版社';
COMMENT ON COLUMN publishers.id IS 'ID';
COMMENT ON COLUMN publishers.name IS '出版社名';
COMMENT ON COLUMN publishers.created_at IS '作成日時';
COMMENT ON COLUMN publishers.updated_at IS '更新日時';
COMMENT ON COLUMN publishers.version_number IS 'バージョン番号';

CREATE TABLE books (
  id INT NOT NULL DEFAULT nextval('books_id_seq'::regclass),
  name VARCHAR(200),
  publisher_id INT NOT NULL,
  isbn VARCHAR(16),
  price NUMBER(10, 4),
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  version_number INT NOT NULL DEFAULT 1
);

COMMENT ON TABLE books IS '書籍';
COMMENT ON COLUMN books.id IS 'ID';
COMMENT ON COLUMN books.name IS '書籍名';
COMMENT ON COLUMN books.publisher_id IS '出版社ID';
COMMENT ON COLUMN books.isbn IS 'ISBNコード';
COMMENT ON COLUMN books.price IS '購入金額';
COMMENT ON COLUMN books.created_at IS '作成日時';
COMMENT ON COLUMN books.updated_at IS '更新日時';
COMMENT ON COLUMN books.version_number IS 'バージョン番号';

CREATE TABLE book_writer_association (
  id INT NOT NULL DEFAULT nextval('book_writer_association_id_seq'::regclass),
  book_id INT NOT NULL,
  writer_id INT NOT NULL,
  writer_role_id INT NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  version_number INT NOT NULL DEFAULT 1
);

COMMENT ON TABLE book_writer_association IS '著者著書関連';
COMMENT ON COLUMN book_writer_association.id IS 'ID';
COMMENT ON COLUMN book_writer_association.book_id IS '蔵書ID';
COMMENT ON COLUMN book_writer_association.writer_id IS '著者ID';
COMMENT ON COLUMN book_writer_association.writer_role_id IS '著者役割ID';
COMMENT ON COLUMN book_writer_association.created_at IS '作成日時';
COMMENT ON COLUMN book_writer_association.updated_at IS '更新日時';
COMMENT ON COLUMN book_writer_association.version_number IS 'バージョン番号';

INSERT INTO writer_role (name) VALUES('著者');
INSERT INTO writer_role (name) VALUES('訳者');
INSERT INTO writer_role (name) VALUES('監修');

INSERT INTO writers (family_name, given_name) VALUES('Richerdson', 'Chris');
INSERT INTO writers (family_name, given_name) VALUES('長尾', '高弘');
INSERT INTO writers (family_name, given_name) VALUES('横澤', '広亨');
