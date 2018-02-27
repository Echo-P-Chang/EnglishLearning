USE [master]
GO
/****** Object:  Database [EnglishLearning]    Script Date: 1/30/2018 7:29:21 AM ******/
CREATE DATABASE [EnglishLearning]
 CONTAINMENT = NONE
 ON  PRIMARY 
( NAME = N'EnglishLearning', FILENAME = N'C:\Program Files\Microsoft SQL Server\MSSQL11.MSSQLSERVER\MSSQL\DATA\EnglishLearning.mdf' , SIZE = 8192KB , MAXSIZE = UNLIMITED, FILEGROWTH = 10%)
 LOG ON 
( NAME = N'EnglishLearning_log', FILENAME = N'C:\Program Files\Microsoft SQL Server\MSSQL11.MSSQLSERVER\MSSQL\DATA\EnglishLearning_log.ldf' , SIZE = 1024KB , MAXSIZE = 2048GB , FILEGROWTH = 10%)
GO

IF (1 = FULLTEXTSERVICEPROPERTY('IsFullTextInstalled'))
begin
EXEC [EnglishLearning].[dbo].[sp_fulltext_database] @action = 'enable'
end
GO
ALTER DATABASE [EnglishLearning] SET ANSI_NULL_DEFAULT OFF 
GO
ALTER DATABASE [EnglishLearning] SET ANSI_NULLS OFF 
GO
ALTER DATABASE [EnglishLearning] SET ANSI_PADDING OFF 
GO
ALTER DATABASE [EnglishLearning] SET ANSI_WARNINGS OFF 
GO
ALTER DATABASE [EnglishLearning] SET ARITHABORT OFF 
GO
ALTER DATABASE [EnglishLearning] SET AUTO_CLOSE OFF 
GO
ALTER DATABASE [EnglishLearning] SET AUTO_SHRINK OFF 
GO
ALTER DATABASE [EnglishLearning] SET AUTO_UPDATE_STATISTICS ON 
GO
ALTER DATABASE [EnglishLearning] SET CURSOR_CLOSE_ON_COMMIT OFF 
GO
ALTER DATABASE [EnglishLearning] SET CURSOR_DEFAULT  GLOBAL 
GO
ALTER DATABASE [EnglishLearning] SET CONCAT_NULL_YIELDS_NULL OFF 
GO
ALTER DATABASE [EnglishLearning] SET NUMERIC_ROUNDABORT OFF 
GO
ALTER DATABASE [EnglishLearning] SET QUOTED_IDENTIFIER OFF 
GO
ALTER DATABASE [EnglishLearning] SET RECURSIVE_TRIGGERS OFF 
GO
ALTER DATABASE [EnglishLearning] SET  DISABLE_BROKER 
GO
ALTER DATABASE [EnglishLearning] SET AUTO_UPDATE_STATISTICS_ASYNC OFF 
GO
ALTER DATABASE [EnglishLearning] SET DATE_CORRELATION_OPTIMIZATION OFF 
GO
ALTER DATABASE [EnglishLearning] SET TRUSTWORTHY OFF 
GO
ALTER DATABASE [EnglishLearning] SET ALLOW_SNAPSHOT_ISOLATION OFF 
GO
ALTER DATABASE [EnglishLearning] SET PARAMETERIZATION SIMPLE 
GO
ALTER DATABASE [EnglishLearning] SET READ_COMMITTED_SNAPSHOT OFF 
GO
ALTER DATABASE [EnglishLearning] SET HONOR_BROKER_PRIORITY OFF 
GO
ALTER DATABASE [EnglishLearning] SET RECOVERY FULL 
GO
ALTER DATABASE [EnglishLearning] SET  MULTI_USER 
GO
ALTER DATABASE [EnglishLearning] SET PAGE_VERIFY CHECKSUM  
GO
ALTER DATABASE [EnglishLearning] SET DB_CHAINING OFF 
GO
ALTER DATABASE [EnglishLearning] SET FILESTREAM( NON_TRANSACTED_ACCESS = OFF ) 
GO
ALTER DATABASE [EnglishLearning] SET TARGET_RECOVERY_TIME = 0 SECONDS 
GO

USE [EnglishLearning]
GO
/****** Object:  Table [dbo].[Answer]    Script Date: 1/30/2018 7:29:21 AM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[Answer](
	[AsnwerID] [int] IDENTITY(1,1) NOT NULL,
	[UserID] [int] NOT NULL,
	[QuestionID] [int] NOT NULL,
	[AnswerBody] [nvarchar](200) NOT NULL,
	[ANswerAttachment] [nvarchar](200) NULL,
 CONSTRAINT [PK_Answer] PRIMARY KEY CLUSTERED 
(
	[AsnwerID] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[Option]    Script Date: 1/30/2018 7:29:22 AM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[Option](
	[OptionID] [int] IDENTITY(1,1) NOT NULL,
	[QuestionID] [int] NOT NULL,
	[OptionAnswer] [nvarchar](100) NULL,
	[OptionAttachment] [nvarchar](200) NULL,
 CONSTRAINT [PK_Option] PRIMARY KEY CLUSTERED 
(
	[OptionID] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[Question]    Script Date: 1/30/2018 7:29:22 AM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[Question](
	[ID] [int] IDENTITY(1,1) NOT NULL,
	[QuestionType] [char](1) NOT NULL,
	[QuestionTitle] [nvarchar](200) NULL,
	[QuestionDescription] [nvarchar](200) NULL,
	[Answer] [nvarchar](200) NULL,
	[QuestionGroup] [int] NULL,
 CONSTRAINT [PK_Question] PRIMARY KEY CLUSTERED 
(
	[ID] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[QuestionType]    Script Date: 1/30/2018 7:29:22 AM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[QuestionType](
	[TypeID] [int] NOT NULL,
	[TypeDescription] [nvarchar](50) NOT NULL,
 CONSTRAINT [PK_QuestionType] PRIMARY KEY CLUSTERED 
(
	[TypeID] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[User]    Script Date: 1/30/2018 7:29:22 AM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[User](
	[UserID] [int] IDENTITY(1,1) NOT NULL,
	[UserName] [nvarchar](50) NOT NULL,
	[UserPassword] [nvarchar](100) NOT NULL,
 CONSTRAINT [PK_User] PRIMARY KEY CLUSTERED 
(
	[UserID] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY]
GO
ALTER TABLE [dbo].[Answer]  WITH CHECK ADD  CONSTRAINT [FK_Answer_Question] FOREIGN KEY([QuestionID])
REFERENCES [dbo].[Question] ([ID])
GO
ALTER TABLE [dbo].[Answer] CHECK CONSTRAINT [FK_Answer_Question]
GO
ALTER TABLE [dbo].[Answer]  WITH CHECK ADD  CONSTRAINT [FK_Answer_User] FOREIGN KEY([UserID])
REFERENCES [dbo].[User] ([UserID])
GO
ALTER TABLE [dbo].[Answer] CHECK CONSTRAINT [FK_Answer_User]
GO
ALTER TABLE [dbo].[Option]  WITH CHECK ADD  CONSTRAINT [FK_Option_Question] FOREIGN KEY([QuestionID])
REFERENCES [dbo].[Question] ([ID])
GO
ALTER TABLE [dbo].[Option] CHECK CONSTRAINT [FK_Option_Question]
GO
/****** Object:  StoredProcedure [dbo].[Insert_Answer]    Script Date: 1/30/2018 7:29:22 AM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE PROCEDURE [dbo].[Insert_Answer] 
    @UserID      INT,
    @QuestionID     INT,
    @AnswerBody    NVARCHAR(200),
    @ANswerAttachment      NVARCHAR(200)
AS
BEGIN

SET NOCOUNT ON;

INSERT INTO [dbo].[Answer]
           ([UserID]
           ,[QuestionID]
           ,[AnswerBody]
           ,[ANswerAttachment])
     VALUES
           (@UserID, @QuestionID, @AnswerBody, @ANswerAttachment)

RETURN @@ROWCOUNT

END
GO
/****** Object:  StoredProcedure [dbo].[Insert_Option]    Script Date: 1/30/2018 7:29:22 AM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO

CREATE PROCEDURE [dbo].[Insert_Option] 
    @QuestionID      INT,
    @OptionAnswer    nvarchar(100),
    @OptionAttachment   nvarchar(200)
AS
BEGIN

SET NOCOUNT ON;

INSERT INTO [dbo].[Option]
           ([QuestionID]
           ,[OptionAnswer]
           ,[OptionAttachment])
     VALUES
           (@QuestionID,@OptionAnswer,@OptionAttachment)

RETURN @@ROWCOUNT

END
GO
/****** Object:  StoredProcedure [dbo].[Insert_Question]    Script Date: 1/30/2018 7:29:22 AM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO


CREATE PROCEDURE [dbo].[Insert_Question] 
    @QuestionType      char(1),
    @QuestionTitle nvarchar(200),
    @QuestionDescription nvarchar(200),
	@Answer nvarchar(200),
	@QuestionGroup int
AS
BEGIN

SET NOCOUNT ON;

INSERT INTO [dbo].[Question]
           ([QuestionType]
           ,[QuestionTitle]
           ,[QuestionDescription]
           ,[Answer]
           ,[QuestionGroup])
     VALUES
           (@QuestionType,@QuestionTitle,@QuestionDescription,@Answer,@QuestionGroup)

RETURN @@ROWCOUNT

END
GO
/****** Object:  StoredProcedure [dbo].[Insert_User]    Script Date: 1/30/2018 7:29:22 AM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO


CREATE PROCEDURE [dbo].[Insert_User] 
    @UserName nvarchar(50),
    @UserPassword nvarchar(100)
AS
BEGIN

SET NOCOUNT ON;

INSERT INTO [dbo].[User]
           ([UserName]
           ,[UserPassword])
     VALUES
           (@UserName,@UserPassword)

RETURN @@ROWCOUNT

END
GO
/****** Object:  StoredProcedure [dbo].[Select_AnswerByUserID]    Script Date: 1/30/2018 7:29:22 AM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE PROCEDURE [dbo].[Select_AnswerByUserID] 
    @UserID      INT
AS
BEGIN

SET NOCOUNT ON;

SELECT * FROM [dbo].[Answer] WHERE [UserID]=@UserID



END
GO
/****** Object:  StoredProcedure [dbo].[Select_OptionByuQuestionID]    Script Date: 1/30/2018 7:29:22 AM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO

CREATE PROCEDURE [dbo].[Select_OptionByuQuestionID] 
    @QuestionID      INT
AS
BEGIN

SET NOCOUNT ON;

SELECT * FROM [dbo].[option] WHERE QuestionID=@QuestionID



END
GO
/****** Object:  StoredProcedure [dbo].[Select_QuestionByQuestionGroup]    Script Date: 1/30/2018 7:29:22 AM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO

CREATE PROCEDURE [dbo].[Select_QuestionByQuestionGroup] 
    @QuestionGroup      INT
AS
BEGIN

SET NOCOUNT ON;

SELECT * FROM [dbo].Question WHERE QuestionGroup=@QuestionGroup



END
GO
/****** Object:  StoredProcedure [dbo].[Select_UserByUserIDAndPassword]    Script Date: 1/30/2018 7:29:22 AM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO

CREATE PROCEDURE [dbo].[Select_UserByUserIDAndPassword] 
    @UserName nvarchar(50),
	@UserPassword nvarchar(100)
AS
BEGIN

SET NOCOUNT ON;

SELECT [UserID]
  FROM [dbo].[User] 
  WHERE [UserName]= @UserName AND [UserPassword]=@UserPassword



END
GO
USE [master]
GO
ALTER DATABASE [EnglishLearning] SET  READ_WRITE 
GO
