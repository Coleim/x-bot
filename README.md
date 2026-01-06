# Twitter Bot (Go) with AI-Generated Content

A smart Twitter bot written in Go that posts engaging AI-generated messages to your Twitter/X account.

## Features

- 🤖 **AI-Generated Tweets** - Uses OpenAI to create engaging content for developers
- 🎯 **Developer-Focused** - Targets software engineers, developers, and game developers
- 🎭 **Mixed Tone** - Alternates between humorous and serious/insightful content
- ⏰ **Cron-Based** - Runs on schedule via Render cron jobs (free tier friendly!)
- 🐳 **Docker-Ready** - Easy deployment with Docker
- 🔐 **Secure** - All credentials managed via environment variables

## Prerequisites

- Go 1.21 or higher (for local development)
- Twitter Developer Account with Read & Write permissions
- Groq API key (FREE - recommended, no credit card!) OR OpenAI API key
- GitHub account (for free automated deployment)

## Setup

### 1. Get API Credentials

Follow the detailed instructions in [KEY.MD](KEY.MD) to:
- Create a Twitter Developer account (free tier works!)
- Create an app with "Read and Write" permissions
- Generate API keys and access tokens
- Get Groq API key (FREE - https://console.groq.com) OR OpenAI API key

### 2. Local Development

```bash
# Install dependencies
go mod download

# Set environment variables
export TWITTER_CONSUMER_KEY="your_consumer_key"
export TWITTER_CONSUMER_SECRET="your_consumer_secret"
export TWITTER_ACCESS_TOKEN="your_access_token"
export TWITTER_ACCESS_SECRET="your_access_secret"

# Add AI for tweet generation (Groq is FREE!)
export AI_API_KEY="gsk_your_groq_key"  # Required - Get from console.groq.com
export AI_PROVIDER="groq"  # or "openai" if you prefer

# Run the bot
go run main.go
```

### 3. Deploy with GitHub Actions (FREE)

1. Push this repository to GitHub
2. Go to your repo → Settings → Secrets and variables → Actions
3. Click "New repository secret" and add each variable:
   - `TWITTER_CONSUMER_KEY`
   - `TWITTER_CONSUMER_SECRET`
   - `TWITTER_ACCESS_TOKEN`
   - `TWITTER_ACCESS_SECRET`
   - `AI_API_KEY`
   - `AI_PROVIDER` (set to `groq`)
4. The workflow runs automatically daily at 12:00 UTC
5. Customize schedule in `.github/workflows/twitter-bot.yml`

The workflow is already configured in the repo!

## Configuration

### Environment Variables

| Variable | Required | Description |
|----------|----------|-------------|
| `TWITTER_CONSUMER_KEY` | Yes | Twitter API Key |
| `TWITTER_CONSUMER_SECRET` | Yes | Twitter API Secret |
| `TWITTER_ACCESS_TOKEN` | Yes | Your Access Token |
| `TWITTER_ACCESS_SECRET` | Yes | Your Access Token Secret |
| `AI_API_KEY` | Yes | Groq or OpenAI API key for AI tweets |
| `AI_PROVIDER` | No | "groq" (free, default) or "openai" |

## Project Structure

```
.
├── main.go           # Main bot application
├── go.mod            # Go dependencies
├── go.sum            # Dependency checksums
├── Dockerfile        # Docker build configuration
├── render.yaml       # Render deployment config
├── KEY.MD            # API key setup instructions
└── README.md         # This file
```

## How It Works

1. Bot authenticates with Twitter API v2 using OAuth1
2. Generates an engaging tweet using AI (Groq or OpenAI)
3. If AI generation fails, bot exits without posting
4. Posts the tweet and exits (perfect for cron jobs)
5. AI creates varied content: humor, insights, tech tips, dev life observations
6. Runs on your schedule via Render cron jobs

## AI Tweet Generation

The bot uses AI to create engaging tweets:
- **Groq (Recommended - FREE)**: Uses Llama 3.3 70B model
  - 14,400 requests/day free tier
  - No credit card required
  - Sign up: https://console.groq.com
- **OpenAI**: Uses GPT-4o-mini
  - ~$0.0001 per tweet
  - $5 free credit initially

**Tweet styles:**
- **Humorous tweets**: Relatable dev struggles, funny observations, coding memes
- **Serious tweets**: Technical insights, career advice, best practices
- **Topics**: Coding, debugging, dev tools, tech trends, game dev, architecture
- **Engagement**: Designed to spark conversation and get likes/retweets

## Building Docker Image Locally

```bash
docker build -t twitter-bot .

docker run -e TWITTER_CONSUMER_KEY="..." \
           -e TWITTER_CONSUMER_SECRET="..." \
           -e TWITTER_ACCESS_TOKEN="..." \
           -e TWITTER_ACCESS_SECRET="..." \
           twitter-bot
```

## Troubleshooting

See [KEY.MD](KEY.MD) for common issues and solutions.

## License

MIT

## Contributing

Feel free to open issues or submit pull requests!
