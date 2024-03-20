'use server'

import {cookies} from 'next/headers'
import { decrypt } from '@/app/lib/data/auth';

export default async function getSessionAction() {
    if (!cookies().get('session')) {
        return null;
    }
    const session = cookies().get('session')?.value;
    if (!session) {
        return null;
    }
    const decryptedSession = await decrypt(session);
    console.log("getSessionAction(): ", decryptedSession)
    return decryptedSession;
}