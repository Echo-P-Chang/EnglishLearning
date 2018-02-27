USE [EnglishLearning]
GO

/****** Object:  Table [dbo].[Tracking]    Script Date: 2/27/2018 8:35:57 AM ******/
SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO

CREATE TABLE [dbo].[Tracking](
	[TrackingID] [int] IDENTITY(1,1) NOT NULL,
	[UserID] [int] NOT NULL,
	[SessionID] [nvarchar](50) NOT NULL,
	[Stage] [nvarchar](50) NOT NULL,
	[Msg] [nvarchar](100) NULL,
	[LogTime] [datetime] NOT NULL,
	[QuestionID] [int] NOT NULL,
	[Answer] [nvarchar](200) NULL,
 CONSTRAINT [PK_Tracking] PRIMARY KEY CLUSTERED 
(
	[TrackingID] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY]
GO

ALTER TABLE [dbo].[Tracking]  WITH CHECK ADD  CONSTRAINT [FK_Tracking_Question] FOREIGN KEY([QuestionID])
REFERENCES [dbo].[Question] ([ID])
GO

ALTER TABLE [dbo].[Tracking] CHECK CONSTRAINT [FK_Tracking_Question]
GO

ALTER TABLE [dbo].[Tracking]  WITH CHECK ADD  CONSTRAINT [FK_Tracking_Tracking] FOREIGN KEY([UserID])
REFERENCES [dbo].[User] ([UserID])
GO

ALTER TABLE [dbo].[Tracking] CHECK CONSTRAINT [FK_Tracking_Tracking]
GO


