# Twitter Bot API Keys and Configuration

## Required Environment Variables

You need to set these environment variables in Render:

### 1. TWITTER_CONSUMER_KEY
- **Description**: Your Twitter App's API Key (Consumer Key)
- **Where to get it**: https://developer.twitter.com/en/portal/dashboard
  1. Go to Twitter Developer Portal
  2. Create a new App (or use existing)
  3. Go to "Keys and tokens" tab
  4. Copy the "API Key"

### 2. TWITTER_CONSUMER_SECRET
- **Description**: Your Twitter App's API Secret Key (Consumer Secret)
- **Where to get it**: Same location as above
  - Copy the "API Key Secret"

### 3. TWITTER_ACCESS_TOKEN
- **Description**: Your account's Access Token
- **Where to get it**: Same "Keys and tokens" tab
  1. Scroll down to "Access Token and Secret"
  2. Click "Generate" if not already generated
  3. Copy the "Access Token"

### 4. TWITTER_ACCESS_SECRET
- **Description**: Your account's Access Token Secret
- **Where to get it**: Same location as above
  - Copy the "Access Token Secret"

### 5. AI_API_KEY (Required) ⭐
- **Description**: API key for AI tweet generation
- **FREE Option - Groq (Recommended)**: 
  - Go to https://console.groq.com
  - Sign up (no credit card required!)
  - Go to "API Keys" section
  - Click "Create API Key"
  - Copy the key (starts with "gsk_...")
  - **Limits**: 14,400 requests/day (more than enough!)
  - **Cost**: 100% FREE forever
  
- **Paid Option - OpenAI**:
  - Go to https://platform.openai.com/api-keys
  - Create account ($5 free credit initially)
  - Generate API key (starts with "sk-...")
  - **Cost**: ~$0.0001 per tweet with gpt-4o-mini
  
- **Without this key**: Bot will use simple fallback messages

### 6. AI_PROVIDER (Optional)
- **Description**: Which AI provider to use
- **Default**: "groq" (free)
- **Options**: "groq" or "openai"
- **Note**: Only needed if you want to use OpenAI instead of Groq



## Setup Instructions

### Step 1: Create Twitter Developer Account
1. Go to https://developer.twitter.com/
2. Sign up for a developer account (Free tier is sufficient)
3. Create a new Project and App

### Step 2: Configure App Permissions **CRITICAL**
1. In your App settings, go to "Settings" tab
2. Click "Edit" next to "User authentication settings"
3. **App permissions**: Select "Read and write" (NOT Read-only!)
4. **Type of App**: Select "Web App, Automated App or Bot"
5. Fill in dummy URLs if needed (e.g., https://example.com)
6. Save changes

### Step 3: Generate Tokens WITH CORRECT PERMISSIONS
1. Go to "Keys and tokens" tab
2. **IMPORTANT**: After changing permissions, you MUST regenerate tokens!
3. Under "Access Token and Secret", click "Regenerate"
4. Confirm regeneration (old tokens will stop working)
5. Copy the new Access Token and Access Token Secret
6. **Verify**: Look for "Created with Read and Write permissions" text

### Common Issue: "oauth1 app permissions" Error
If you see this error, your access tokens were generated with Read-only permissions.
**Solution**: 
1. Go to Settings → User authentication settings
2. Change to "Read and write"
3. Save
4. Go to Keys and tokens
5. **Regenerate** Access Token and Secret (don't reuse old ones!)
6. Update your environment variables with the new tokens

### Step 4: Deploy to Render
1. Push this code to a GitHub repository
2. Go to https://render.com/
3. Create a new "Web Service"
4. Connect your GitHub repository
5. Render will auto-detect the Dockerfile
6. Add the environment variables from above
7. Deploy!

## Troubleshooting

### Error: "Read-only application cannot POST"
- Your app permissions are set to "Read-only"
- Change to "Read and Write" in Twitter Developer Portal
- Regenerate your Access Token and Secret
- Update the tokens in Render

### Error: "Could not authenticate you"
- Check that all 4 credentials are correct
- Make sure there are no extra spaces
- Verify the tokens haven't been revoked

### Error: "Missing required environment variables"
- All 4 Twitter credentials must be set in Render
- Go to your service's "Environment" tab
- Add missing variables

## Bot Behavior

- **On Start**: Posts the TWEET_MESSAGE once
- **Periodic**: Posts PERIODIC_MESSAGE every 24 hours with a timestamp
- **Runs continuously**: The bot stays running on Render

## Cost

- Render: Free tier available (Starter plan)
- Twitter API: Free tier available (Elevated access)

## Security Notes

⚠️ **NEVER commit your actual API keys to git!**
- All credentials should be set as environment variables in Render
- Never hardcode them in the source code
- Keep this KEY.MD file as a reference only
