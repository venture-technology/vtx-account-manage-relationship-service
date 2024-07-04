-- Tabela de school-drivers
CREATE TABLE IF NOT EXISTS schools_drivers (
    record SERIAL PRIMARY KEY,
    name_school VARCHAR(100) NOT NULL,
    cnpj_school VARCHAR(14),
    email_school VARCHAR(100) NOT NULL,
    name_driver VARCHAR(100) NOT NULL,
    cnh_driver VARCHAR(14),
    email_driver VARCHAR(100) NOT NULL
);

-- Tabela de account-manager
CREATE TABLE IF NOT EXISTS partners (
    record SERIAL PRIMARY KEY,
    name_school VARCHAR(100) NOT NULL,
    cnpj_school VARCHAR(14) NOT NULL,
    email_school VARCHAR(100) NOT NULL,
    name_driver VARCHAR(100) NOT NULL,
    cnh_driver VARCHAR(14) NOT NULL,
    cpf_responsible VARCHAR(11) NOT NULL,
    name_responsible VARCHAR(100) NOT NULL,
    email_responsible VARCHAR(100) NOT NULL,
    name_child VARCHAR(50) NOT NULL,
    rg_child VARCHAR(10) NOT NULL
);