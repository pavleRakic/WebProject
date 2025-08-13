IF OBJECT_ID('dbo.Users', 'U') IS NULL
BEGIN
    CREATE TABLE dbo.Users
    (
        idUser INT IDENTITY(1,1) NOT NULL
    );
END