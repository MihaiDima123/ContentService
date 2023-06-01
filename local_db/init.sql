create table tenant (
    id SERIAL primary key,
    tenant_ref varchar(50) unique,
    channel varchar(20),
    language varchar(5),
    created_at DATE NOT NULL DEFAULT CURRENT_DATE,
    updated_at DATE
);

create table configuration (
    id SERIAL primary key,
    Configuration_Name varchar(30) unique,
    Configuration_Description varchar(50)
);

create table tenant_configuration (
    id SERIAL primary key,
    tenant_id int,
    configuration_id int,
    created_at DATE NOT NULL DEFAULT CURRENT_DATE,
    updated_at DATE,
    constraint fk_tenant_configuration_tenant foreign key (tenant_id) references tenant(id),
    constraint fk_tenant_configuration_configuration foreign key (configuration_id) references configuration(id)
);

create table category(
    id SERIAL primary key,
    name varchar(50),
    tenant_id int not null,
    constraint fk_user_content_tenant foreign key (tenant_id) references tenant(id),
    constraint uk_user_content_name_tenant unique (tenant_id, name)
);

-- Template
create table post_template(
    id SERIAL primary key,
    name varchar(50) not null,
    tenant_id int not null,
    constraint fk_user_content_tenant foreign key (tenant_id) references tenant(id)
);

create table post_template_slot(
    id SERIAL primary key,
    name varchar(50),
    parent_id int,
    tenant_id int not null,
    constraint fk_post_template_slot_parent foreign key (parent_id) references post_template_slot(id),
    constraint fk_post_template_slot_tenant foreign key (tenant_id) references tenant(id)
);

create table post_template_slot_property_type(
    id SERIAL primary key,
    name varchar(20) not null,
    tenant_id int not null,
    constraint fk_post_template_slot_property_type_tenant foreign key (tenant_id) references tenant(id),
    constraint fk_post_template_slot_property_type_name_tenant unique (tenant_id, name)
);

create table post_template_slot_property(
    id SERIAL primary key,
    name varchar(50) not null,
    value varchar(256) not null,
    property_type_id int not null,
    post_template_slot_id int not null,
    tenant_id int not null,
    constraint fk_post_template_slot_property_post_template_slot foreign key (post_template_slot_id) references post_template_slot(id),
    constraint fk_post_template_slot_property_post_template_slot_property_type foreign key (property_type_id) references post_template_slot_property_type (id),
    constraint fk_post_template_slot_property_tenant foreign key (tenant_id) references tenant(id)
);

-- Post
create table post(
    id SERIAL primary key,
    user_id int not null,
    user_email varchar(255),
    category_id int,
    post_template_id int not null,
    tenant_id int not null,
    constraint fk_user_content_tenant foreign key (tenant_id) references tenant(id),
    constraint fk_user_content_category foreign key (category_id) references category(id),
    constraint fk_user_content_template foreign key (post_template_id) references post_template(id)
);

create table post_content(
     id SERIAL primary key,
     content text,
     post_id int not null,
     post_template_slot_id int not null,
     tenant_id int,
     constraint fk_user_content_post_id foreign key (post_id) references post(id),
     constraint fk_user_content_tenant foreign key (tenant_id) references tenant(id),
     constraint fk_user_content_slot foreign key (post_template_slot_id) references post_template_slot(id)
);

insert into tenant (tenant_ref, channel, language) values ('blogs-ro_RO', 'blogs', 'ro_RO');
