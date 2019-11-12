import React, { useState, useEffect } from 'react'
import PropTypes from 'prop-types';
import {
    makeStyles,
    Dialog,
    DialogTitle,
    Typography,
    DialogContent,
    DialogContentText,
    TextField,
    FormControl,
    Button,
    DialogActions
} from '@material-ui/core'

const useStyles = makeStyles(theme => ({
    titleRoot: {
        backgroundColor: theme.palette.primary.main,
        padding: theme.spacing(2),
    },
    container: {
        padding: theme.spacing(2),
    },
    formControl: {
        width: '100%',
        marginTop: theme.spacing(1),
    },
    fields: {
        marginTop: theme.spacing(1),
    },
    titleText: {
        color: 'white',
    }
}))

const WorkshopForm = (props) => {
    const classes = useStyles()

    const { open, handleClose, handleSave, ws } = props;

    const [workshop, setWorkshop] = useState(ws)

    const handleOnChange = (event) => {
        event.persist();
        setWorkshop(workshop => ({ ...workshop, [event.target.name]: event.target.value }));
    }

    const handleSaveButton = () => {
        handleSave(workshop)
        handleClose()
    }

    useEffect(() => {
        setWorkshop(ws)
    }, [ws]);

    return (
        <Dialog open={open} onClose={handleClose} aria-labelledby="form-dialog-title">
            <DialogTitle disableTypography className={classes.titleRoot} id="form-dialog-title">
                <Typography className={classes.titleText} variant="h5">Workshop</Typography>
            </DialogTitle>
            <DialogContent classes={{ root: classes.container }} dividers>
                <DialogContentText>
                    Para cadastrar ou alterar um Workshop, por favor preencha os campos abaixo
                </DialogContentText>
                <FormControl className={classes.formControl}>
                    <TextField fullWidth
                        label="Nome do workshop" margin="dense" name="nome"
                        onChange={handleOnChange} required
                        value={workshop.nome} variant="outlined" />
                </FormControl>
                <FormControl className={classes.formControl}>
                    <TextField fullWidth
                        label="Palestrante" margin="dense" name="palestrante"
                        onChange={handleOnChange} required
                        value={workshop.palestrante} variant="outlined" />
                </FormControl>
            </DialogContent>
            <DialogActions>
                <Button onClick={handleSaveButton} variant="contained" color="primary">
                    Salvar
                </Button>
                <Button onClick={handleClose} variant="contained" color="primary">
                    Cancelar
              </Button>
            </DialogActions>
        </Dialog>
    )
}

WorkshopForm.propTypes = {
    ws: PropTypes.object,
    open: PropTypes.bool,
    handleClose: PropTypes.func,
    handleSave: PropTypes.func,
};

export default WorkshopForm;