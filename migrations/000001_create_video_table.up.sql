CREATE TYPE video_status AS ENUM(
  'uploaded',
  'processing',
  'queued',
  'ready',
  'failed'
);
CREATE TABLE videos(
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  title TEXT NOT NULL,
  orignal_url TEXT NOT NULL,
  video_360_url TEXT NOT NULL,
  video_480_url TEXT NOT NULL,
  video_720_url TEXT NOT NULL,
  video_1080_url TEXT NOT NULL,
  status video_status NOT NULL DEFAULT 'uploaded',
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
