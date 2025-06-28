import asyncio
import os
from typing import AsyncGenerator

from telethon import TelegramClient
from telethon.tl.types import Message, MessageMediaDocument, MessageMediaPhoto

API_ID = "20887348"
API_HASH = "6621da0e5ba2c3eba558269d38e8afa9"
CHANNEL_USERNAME = "@statistikbottesting"
TOKEN = "7618324350:AAG7bQ8gQJ0AVnShATl1X3Vhl5q_q6ROvVw"

PHONE = "+79931437593"
PASS = "z1v4m7ALG"
MY_CHAT_ID = "1172084194"
MY_USERNAME = "@max1t1a"


async def find_all_messages() -> AsyncGenerator[Message, None]:
    async with await TelegramClient("client", API_ID, API_HASH).start(phone=PHONE, password=PASS) as client:
        channel = await client.get_entity(CHANNEL_USERNAME)
        async for message in client.iter_messages(channel):
            yield message


async def send_media_for_me():
    async with await TelegramClient("bot", API_ID, API_HASH).start(bot_token=TOKEN) as bot:

        async for message in find_all_messages():
            if isinstance(message.media, (MessageMediaPhoto, MessageMediaDocument)):
                try:
                    media_data = await message.download_media()
                    await bot.send_file(
                        entity=MY_USERNAME,
                        file=media_data,
                        caption=f"Медиа из канала {CHANNEL_USERNAME}\n\n{message.text}"[:1024],
                        force_document=False,
                    )
                    print(f"Отправлено медиа из сообщения {message.id}")

                    os.remove(media_data)

                    await asyncio.sleep(1)

                except Exception as e:
                    print(f"Ошибка при отправке {message.id}: {str(e)}")


async def main():
    await send_media_for_me()


if __name__ == "__main__":
    asyncio.run(main())
