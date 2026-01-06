# Get Your FREE Groq API Key 🚀

Groq provides 100% FREE AI inference with no credit card required!

## Step-by-Step Instructions

### 1. Sign Up
- Go to https://console.groq.com
- Click "Sign Up" or "Get Started"
- Use Google, GitHub, or email to create account
- ✅ **NO credit card required!**

### 2. Create API Key
- Once logged in, look for "API Keys" in the left sidebar
- Click "Create API Key"
- Give it a name (e.g., "Twitter Bot")
- Click "Create"
- **Copy the key immediately** (starts with `gsk_...`)
- You won't be able to see it again!

### 3. Use the Key
Set it as an environment variable:
```bash
export AI_API_KEY="gsk_your_key_here"
export AI_PROVIDER="groq"
```

Or add to Render:
- Environment variable: `AI_API_KEY`
- Value: `gsk_your_key_here`

## Free Tier Limits

✅ **14,400 requests per day**
- That's 600 requests per hour!
- Perfect for 1 tweet per day (you only need 1!)
- No expiration, no credit card, forever free

## Models Available

The bot uses **Llama 3.3 70B Versatile** - one of the best open-source models:
- Fast inference
- High quality outputs
- Great for creative content

## Troubleshooting

### "Invalid API key" error
- Make sure you copied the entire key (starts with `gsk_`)
- Check for extra spaces
- Try regenerating a new key

### "Rate limit exceeded"
- Free tier: 14,400 requests/day
- For 1 tweet/day bot, you'll never hit this!

## Alternative: OpenAI

If you prefer OpenAI (costs money):
1. Go to https://platform.openai.com/api-keys
2. Create account ($5 free credit)
3. Generate API key (starts with `sk-...`)
4. Set `AI_API_KEY="sk_your_key"` and `AI_PROVIDER="openai"`

But Groq is FREE and works great! 🎉
