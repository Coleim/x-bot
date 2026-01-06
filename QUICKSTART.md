# Quick Start - Twitter Bot with FREE AI

## Get Started in 5 Minutes! 🚀

### 1. Get Groq API Key (100% FREE)
```
1. Visit: https://console.groq.com
2. Sign up (no credit card!)
3. Create API Key
4. Copy key (starts with gsk_...)
```

### 2. Set Environment Variables
```bash
# Twitter credentials (from developer.twitter.com)
export TWITTER_CONSUMER_KEY="your_key"
export TWITTER_CONSUMER_SECRET="your_secret"
export TWITTER_ACCESS_TOKEN="your_token"
export TWITTER_ACCESS_SECRET="your_token_secret"

# Groq AI (FREE from console.groq.com)
export AI_API_KEY="gsk_your_groq_key"
export AI_PROVIDER="groq"
```

### 3. Run!
```bash
go run main.go
```

## Deploy with GitHub Actions (100% FREE)

### 1. Push to GitHub
```bash
git init
git add .
git commit -m "Twitter bot with AI"
git branch -M main
git remote add origin YOUR_REPO_URL
git push -u origin main
```

### 2. Add Secrets to GitHub
- Go to your repo on GitHub
- Settings → Secrets and variables → Actions
- Click "New repository secret"
- Add each secret:
  - `TWITTER_CONSUMER_KEY`
  - `TWITTER_CONSUMER_SECRET`
  - `TWITTER_ACCESS_TOKEN`
  - `TWITTER_ACCESS_SECRET`
  - `AI_API_KEY`
  - `AI_PROVIDER` (value: `groq`)

### 3. That's It!
- Bot runs automatically daily at 12:00 UTC
- Check Actions tab to see runs
- Customize schedule in `.github/workflows/twitter-bot.yml`

## What You Get

✅ AI-generated tweets tailored for developers
✅ Mix of humor and technical insights
✅ Automatic posting on schedule
✅ 100% FREE (Groq + Twitter + Render free tiers)
✅ 14,400 tweets/day possible (you only need 1!)

## Cost Breakdown

| Service | Cost | Limit |
|---------|------|-------|
| Twitter API | FREE | Basic features |
| Groq AI | FREE | 14,400 requests/day |
| GitHub Actions | FREE | 2,000 min/month |
| **TOTAL** | **$0/month** | **Perfect for daily tweets** |

## Schedule Examples

Edit `.github/workflows/twitter-bot.yml` cron field:

```yaml
cron: '0 12 * * *'   # Daily at noon UTC
cron: '0 9,17 * * *' # Twice daily: 9am & 5pm UTC
cron: '0 */6 * * *'  # Every 6 hours
cron: '0 8 * * 1-5'  # Weekdays at 8am UTC
```

## Troubleshooting

### Twitter "oauth1 permissions" error
→ Regenerate tokens with "Read and Write" permissions

### "Invalid API key" 
→ Check you copied the full Groq key (starts with `gsk_`)

### "Duplicate content"
→ Twitter blocks identical tweets. Wait or change message.

## Next Steps

📖 Full docs: [README.md](README.md)
🔑 API setup: [KEY.MD](KEY.MD) & [GROQ_SETUP.md](GROQ_SETUP.md)
