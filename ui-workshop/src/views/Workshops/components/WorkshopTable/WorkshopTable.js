import React, { useState } from 'react';
import clsx from 'clsx';
import PropTypes from 'prop-types';
import moment from 'moment';
import PerfectScrollbar from 'react-perfect-scrollbar';
import { makeStyles } from '@material-ui/styles';
import {
  Card,
  CardActions,
  CardContent,
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableRow,
  Typography,
  IconButton,
  Modal,
  Fade,
  Backdrop,
  CardHeader,
  Button
} from '@material-ui/core';
import { Edit, Delete } from '@material-ui/icons';


const useStyles = makeStyles(theme => ({
  root: {},
  content: {
    padding: 0
  },
  inner: {
    minWidth: 1050
  },
  nameContainer: {
    display: 'flex',
    alignItems: 'center'
  },
  avatar: {
    marginRight: theme.spacing(2)
  },
  actions: {
    justifyContent: 'flex-end'
  },
  modal: {
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'center',
  }
}));

const WorkshopTable = props => {
  const { className, workshops, handleDelete, handleEdit, ...rest } = props;

  const [open, setOpen] = useState(false)
  const [workshop, setWorkshop] = useState({})

  const classes = useStyles();

  const handleOpenDelete = (_workshop) => {
    setWorkshop(_workshop)
    setOpen(true)
  }

  const handleDeleteYes = () => {
    setOpen(false)
    handleDelete(workshop)
  }

  const handleCloseDelete = () => {
    setOpen(false)
  }

  return (
    <Card
      {...rest}
      className={clsx(classes.root, className)}
    >
      <CardContent className={classes.content}>
        <PerfectScrollbar>
          <div className={classes.inner}>
            <Table>
              <TableHead>
                <TableRow>
                  <TableCell>ID</TableCell>
                  <TableCell>Nome</TableCell>
                  <TableCell>Palestrante</TableCell>
                  <TableCell>Ações</TableCell>
                </TableRow>
              </TableHead>
              <TableBody>
                {workshops.map(workshop => (
                  <TableRow className={classes.tableRow} hover key={workshop.id}>
                    <TableCell>{workshop.id}</TableCell>
                    <TableCell>
                      <Typography variant="body1">{workshop.nome}</Typography>
                    </TableCell>
                    <TableCell>{workshop.palestrante}</TableCell>
                    <TableCell>
                      <IconButton aria-label="edit" className={classes.margin} onClick={() => handleEdit(workshop)}>
                        <Edit />
                      </IconButton>
                      <IconButton aria-label="delete" className={classes.margin} onClick={() => handleOpenDelete(workshop)}>
                        <Delete />
                      </IconButton>
                    </TableCell>
                  </TableRow>
                ))}
              </TableBody>
            </Table>
          </div>
        </PerfectScrollbar>
      </CardContent>
      <Modal
        aria-labelledby="transition-modal-title"
        aria-describedby="transition-modal-description"
        className={classes.modal}
        open={open}
        onClose={handleCloseDelete}
        closeAfterTransition
        BackdropComponent={Backdrop}
        BackdropProps={{
          timeout: 500,
        }}
      >
        <Fade in={open}>
          <Card {...rest}
            className={clsx(classes.root, className)}>
            <CardHeader
              title="Confirmação" />
            <CardContent>
              <Typography variant="body2" color="textSecondary" component="p">
                Confirma a exclusão do workshop?
            </Typography>
            </CardContent>
            <CardActions>
              <Button onClick={handleDeleteYes}>Sim</Button>
              <Button onClick={handleCloseDelete}>Não</Button>
            </CardActions>
          </Card>

        </Fade>
      </Modal>
    </Card>
  );
};

WorkshopTable.propTypes = {
  className: PropTypes.string,
  workshops: PropTypes.array.isRequired,
  handleEdit: PropTypes.func,
  handleDelete: PropTypes.func
};

export default WorkshopTable;
