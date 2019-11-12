let enqueueSnackbar = undefined;
const options = {
    variant: 'default',
    anchorOrigin: {
        vertical: 'bottom',
        horizontal: 'center'
    }
}

const setEnqueueSnackbar = (snackbar) => {
    console.log(snackbar)
    enqueueSnackbar = snackbar;
}

const showMessageSuccess = (message) => {
    options.variant = 'success';
    enqueueSnackbar(message, options);
}

const showMessageError = (message) => {
    options.variant = 'error';
    enqueueSnackbar(message, options);
}

export {
    setEnqueueSnackbar, showMessageSuccess, showMessageError
}

