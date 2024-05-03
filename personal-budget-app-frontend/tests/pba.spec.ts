import { test, expect } from '@playwright/test';

test('has title', async ({ page }) => {
    await page.goto('localhost:3000/');

    // Expect a title "to contain" a substring.
    await expect(page).toHaveTitle(/Personal Budget App/);
});

test('sign in works', async ({ page }) => {
    await page.goto('localhost:3000/sign-in');
    await page.fill('input[name="email"]', 'test@test.com');
    await page.fill('input[name="password"]', 'testpassword');
    await page.click('button[type="submit"]');
    await expect(page).toHaveURL('http://localhost:3000/dashboard');
});

test('sign in fails', async ({ page }) => {
    await page.goto('localhost:3000/sign-in');
    await page.fill('input[name="email"]', 'test@test.com');
    await page.fill('input[name="password"]', 'wrongpassword');
    await page.click('button[type="submit"]');
    await expect(page.locator('#confirmation-message')).toHaveText('username and/or password incorrect');
});
