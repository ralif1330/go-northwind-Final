CREATE TABLE
    public.categories (
        category_id smallint NOT NULL,
        category_name character varying(15),
        description text,
        picture bytea
    );

ALTER TABLE public.categories OWNER TO postgres;

--

-- Name: customers; Type: TABLE; Schema: public; Owner: postgres

--

CREATE TABLE
    public.customers (
        customer_id character varying(50) NOT NULL,
        company_name character varying(40),
        contact_name character varying(30),
        contact_title character varying(30),
        address character varying(60),
        city character varying(15),
        region character varying(15),
        postal_code character varying(10),
        country character varying(15),
        phone character varying(24),
        fax character varying(24)
    );

ALTER TABLE public.customers OWNER TO postgres;

--

-- Name: employees; Type: TABLE; Schema: public; Owner: postgres

--

CREATE TABLE
    public.employees (
        employee_id smallint NOT NULL,
        last_name character varying(20),
        first_name character varying(10),
        title character varying(30),
        title_of_courtesy character varying(25),
        birth_date date,
        hire_date date,
        address character varying(60),
        city character varying(15),
        region character varying(15),
        postal_code character varying(10),
        country character varying(15),
        home_phone character varying(24),
        extension character varying(4),
        photo bytea,
        notes text,
        reports_to smallint,
        photo_path character varying(255)
    );

ALTER TABLE public.employees OWNER TO postgres;

--

-- Name: order_detail; Type: TABLE; Schema: public; Owner: postgres

--

CREATE TABLE
    public.order_detail (
        order_id smallint NOT NULL,
        product_id smallint NOT NULL,
        unit_price real,
        quantity smallint,
        discount real
    );

ALTER TABLE public.order_detail OWNER TO postgres;

--

-- Name: orders; Type: TABLE; Schema: public; Owner: postgres

--

CREATE TABLE
    public.orders (
        order_id smallint NOT NULL,
        order_date date,
        required_date date,
        shipped_date date,
        freight real,
        ship_name character varying(40),
        ship_address character varying(60),
        ship_city character varying(15),
        ship_region character varying(15),
        ship_postal_code character varying(10),
        ship_country character varying(15),
        employee_id smallint,
        customer_id character varying(100),
        shipper_id smallint
    );

ALTER TABLE public.orders OWNER TO postgres;

--

-- Name: products; Type: TABLE; Schema: public; Owner: postgres

--

CREATE TABLE
    public.products (
        product_id smallint NOT NULL,
        product_name character varying(40),
        quantity_per_unit character varying(20),
        unit_price real,
        unit_in_stock smallint,
        unit_in_order smallint,
        reorder_level smallint,
        discontinued integer,
        supplier_id smallint,
        category_id smallint
    );

ALTER TABLE public.products OWNER TO postgres;

--

-- Name: shippers; Type: TABLE; Schema: public; Owner: postgres

--

CREATE TABLE
    public.shippers (
        shipper_id smallint NOT NULL,
        company_name character varying(40),
        phone character varying(24)
    );

ALTER TABLE public.shippers OWNER TO postgres;

--

-- Name: supplier; Type: TABLE; Schema: public; Owner: postgres

--

CREATE TABLE
    public.supplier (
        supplier_id smallint NOT NULL,
        company_name character varying(40),
        contact_name character varying(30),
        contact_title character varying(30),
        address character varying(60),
        city character varying(15),
        region character varying(15),
        postal_code character varying(10),
        country character varying(15),
        phone character varying(24),
        fax character varying(24),
        homepage text
    );

ALTER TABLE public.supplier OWNER TO postgres;