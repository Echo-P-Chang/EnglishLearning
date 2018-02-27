USE [EnglishLearning]
GO

/****** Object:  Table [dbo].[Question]    Script Date: 1/30/2018 6:56:06 AM ******/
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
	[QuestionGroup] [INT] NULL,
 CONSTRAINT [PK_Question] PRIMARY KEY CLUSTERED 
(
	[ID] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY]
GO


