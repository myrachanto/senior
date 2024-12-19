If in Golang we handle errors immediately when encountering them why is the need for `error.Is()` and `error.As()`


Hi everyone 🤗 , I hope your holidays are fabulous!


so,


✅ `errors.Is`:

 Checks if an error is of a specific type or wraps another error of that type.


✅ `errors.As`:

 Unwraps and assigns an error to a target variable if the error (or a wrapped error) matches the target type.


this is of course made possible given that we can wrap errors since go 1.13


like

`err := fmt.Errorf("failed to open file: %w", os.ErrNotExist)`


so why the error.Is and error.As?


🙏 Do Immediate Error Handling: 

If you can handle the error at the point of occurrence, do so.


but that is not always the case. In the screenshot below in a parent function, we need to know the exact reason why an error occurred.

 

in our case, we are trying to copy the contents of one file to another and an error occurs. it can be


1️⃣ source doesn't exist

2️⃣ destination doesn't exist 

3️⃣ we don't have permission to read or write the file

4️⃣ we got the file path wrong etc


These are just a few ways we could encounter an error.


with error.Is and error.As we can figure out what happened and what the next steps should be.


so basically 


🔆 errors.Is and errors.As are tools to figure out what exactly happened, especially when dealing with filesystem operations or similar scenarios where errors could have multiple causes. They allow you to handle specific cases programmatically, even after adding context or wrapping errors.


happy holidays everyone!

#go #golang #golangtips #senior #backend #seniorgolang #golangdeveloper #errorhandling