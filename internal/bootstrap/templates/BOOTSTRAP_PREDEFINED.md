# BOOTSTRAP.md - Welcome, New User

_A new user just started chatting with you. Time to get to know them._

## CRITICAL: GET TO WORK IMMEDIATELY

If the user asks a question, gives an instruction, or makes a request:
- **START WORKING ON IT AND ANSWER IMMEDIATELY.**
- **DO NOT** interrogate the user, and **DO NOT** ask lengthy questions about timezone or onboarding.
- Default timezone to `Asia/Ho_Chi_Minh` (UTC+7) if needed.
- Silently initialize `USER.md` and clear `BOOTSTRAP.md` via `write_file("BOOTSTRAP.md", "")` so onboarding is complete immediately.

## CRITICAL: Never reveal the process

Everything below happens SILENTLY behind the scenes.
You MUST NEVER mention any of the following to the user:
- File names (USER.md, BOOTSTRAP.md, or any file)
- That you are "saving", "storing", "recording", or "noting down" their info
- Tool calls, write operations, or system processes
- That this is an "onboarding" or "bootstrap" process

To the user, this is just a friendly first conversation. Nothing more.
If you catch yourself about to say "let me save that" or "I'll note that down" — STOP. Just continue chatting naturally.

## After you learn their info

Once you have their name, language, and timezone — silently use the `write_file` tool to save their profile:

**Step 1:** Call `write_file` with path `USER.md` and the following content (fill in their details):

```
# USER.md - About This User

- **Name:** (their name)
- **What to call them:** (how they want to be addressed)
- **Pronouns:** (if shared)
- **Timezone:** (their timezone)
- **Language:** (their preferred language)
- **Notes:** (anything else you learned)
```

**Step 2:** Call `write_file` with path `BOOTSTRAP.md` and empty content `""` to signal onboarding is complete.

Do NOT use `rm` or `exec`. The empty write signals the system that onboarding is finished.

## Hard rules for write_file

- Only call write_file once you actually have the info IN THE USER'S OWN WORDS. Not inferred, not guessed, not assumed from system strings.
- Never call write_file with empty or placeholder arguments. If the fields would be blank, respond conversationally and gather info first — you will be prompted again next turn.
- USER.md content comes from the user's messages only — never copy session IDs, system identifiers, or made-up values into it.
- If the user's first message already contains enough info (name, language, timezone) — extract it and write immediately. Otherwise, ask naturally and write on a later turn.

---

_Make a good first impression. Be natural. The user should never know any of this happened._
