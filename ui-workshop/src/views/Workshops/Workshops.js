import React, { useState, useEffect } from 'react';
import { makeStyles } from '@material-ui/styles';
import { WorkshopForm, WorkshopToolbar, WorkshopTable } from './components';
import { getWorkshops, saveWorkshop, updateWorkshop, deleteWorkshop } from 'helpers/api';
import { setEnqueueSnackbar, showMessageError, showMessageSuccess } from 'helpers/snackbar'
import { withSnackbar, useSnackbar } from 'notistack';

const useStyles = makeStyles(theme => ({
    root: {
        padding: theme.spacing(3)
    },
    content: {
        marginTop: theme.spacing(2)
    }
}));

const Workshops = (props) => {
    const classes = useStyles()
    const { enqueueSnackbar } = props
    const [open, setOpen] = useState(false)
    const [workshop, setWorkshop] = useState({})
    const [workshops, setWorkshops] = useState([])


    const handleOpen = () => () => {
        setWorkshop({})
        setOpen(true)
    }

    const handleEdit = (_workshop) => {
        setWorkshop(_workshop)
        setOpen(true)
    }

    const handleDelete = (_workshop) => {
        deleteWorkshop(_workshop.id, (result) => {
            handleGetWorkshops()
            showMessageSuccess(result.message)
        }, (error) => {
            showMessageError(error)
        })
    }

    const handleSave = (_workshop) => {
        if (_workshop.id) {
            updateWorkshop(_workshop, (result) => {
                handleGetWorkshops()
                showMessageSuccess(result.message)
            }, (error) => {
                showMessageError(error)
            })
        } else {
            saveWorkshop(_workshop, (result) => {
                handleGetWorkshops()
                showMessageSuccess(result.message)
            }, (error) => {
                showMessageError(error)
            })
        }
    }

    const handleClose = (update) => {
        setOpen(false)
        if (update === 'atualizar') {
            handleGetWorkshops()
        }
    }

    const handleGetWorkshops = () => {
        getWorkshops((ws) => {
            setWorkshops(ws)
        }, (error) => {
            showMessageError(error)
        })
    }

    useEffect(() => {
        setEnqueueSnackbar(enqueueSnackbar)
        handleGetWorkshops()
    }, [])

    return (
        <div className={classes.root}>
            <WorkshopForm open={open} ws={workshop} handleClose={handleClose} handleSave={handleSave} />
            <WorkshopToolbar handleOpen={handleOpen} />
            <div className={classes.content}>
                <WorkshopTable workshops={workshops} handleEdit={handleEdit} handleDelete={handleDelete} />
            </div>
        </div>
    )
}

export default withSnackbar(Workshops);