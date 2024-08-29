CREATE TABLE users (
    id bigserial PRIMARY KEY,
    username text NOT NULL,
    email text NOT NULL,
    first_name text NOT NULL,
    last_name text,
    password_hash text,
    blocked bool DEFAULT FALSE,
    active bool DEFAULT FALSE,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);
CREATE UNIQUE INDEX idx_uniq_users_usersname ON public.users USING btree (username);
COMMENT ON TABLE public.users IS 'Пользователи';
COMMENT ON COLUMN public.users.id IS 'Идентификатор пользователя';
COMMENT ON COLUMN public.users.username IS 'Логин пользователя';
COMMENT ON COLUMN public.users.email IS 'Электронная почта';
COMMENT ON COLUMN public.users.first_name IS 'Имя пользователя';
COMMENT ON COLUMN public.users.last_name IS 'Фамилия пользователя';
COMMENT ON COLUMN public.users.blocked IS 'Учётная запись заблокирована';
COMMENT ON COLUMN public.users.active IS 'Учётная запись активна';
COMMENT ON COLUMN public.users.created_at IS 'Дата создания';
COMMENT ON COLUMN public.users.updated_at IS 'Дата обновления';
COMMENT ON COLUMN public.users.deleted_at IS 'Дата удаления';

--bun:split

CREATE TABLE roles (
    id serial PRIMARY KEY,
    name text NOT NULL,
    description text,
    type text,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);
COMMENT ON TABLE public.roles IS 'Роли';
COMMENT ON COLUMN public.roles.id IS 'Идентификатор роли';
COMMENT ON COLUMN public.roles.name IS 'Имя роли';
COMMENT ON COLUMN public.roles.description IS 'Описание роли';
COMMENT ON COLUMN public.roles.type IS 'Тип роли';

--bun:split

INSERT INTO roles (name, description, type, created_at, updated_at)
VALUES
    ('Администратор', 'Администратор сайта', 'admin', NOW(), NOW()),
    ('Пользователь', 'Обычный пользователь', 'user', NOW(), NOW());

--bun:split

CREATE TABLE public.users_to_roles (
    user_id INT8 NOT NULL,
    role_id INT4 NOT NULL,
    CONSTRAINT pk_users_to_roles PRIMARY KEY (user_id, role_id),
    CONSTRAINT fk_users_to_roles_user_id FOREIGN KEY (user_id) REFERENCES public.users(id),
    CONSTRAINT fk_users_to_roles_role_id FOREIGN KEY (role_id) REFERENCES public.roles(id)
);
CREATE INDEX idx_users_to_roles_user_id ON public.users_to_roles USING btree (user_id);
CREATE INDEX idx_users_to_roles_role_id ON public.users_to_roles USING btree (role_id);
