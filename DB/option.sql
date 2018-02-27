USE [EnglishLearning]
GO

/****** Object:  Table [dbo].[Option]    Script Date: 1/30/2018 6:25:08 AM ******/
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

ALTER TABLE [dbo].[Option]  WITH CHECK ADD  CONSTRAINT [FK_Option_Question] FOREIGN KEY([QuestionID])
REFERENCES [dbo].[Question] ([ID])
GO

ALTER TABLE [dbo].[Option] CHECK CONSTRAINT [FK_Option_Question]
GO


