Docker for containerization

PGRF stack - Postgres, Go, React, Fiber

Backend:
    Go on the backend server and API
    Fiber for Express.js-like web framework/server
    CORS let browser to enable request to local ports on same machine 
    GORM as object relational map
    Fasthttp as HTTP engine
    Air for hot-loading
    Authentication:
        jwt with http-only cookie (prone to csrf attacks)
        vs. token based auth (prone to xss attacks)
        best: short lived access token and refresh token https://www.youtube.com/watch?v=rT20ylRLm5U&t=17s

Frontend:
    Yarn for package/project manager. (npm alternative)
    Vite for front end tooling, local hosting & template react-ts
    React for front end UI interface library
    Mantine for React components
    SWR for data fetching using React Hooks